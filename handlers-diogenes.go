package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/lucianonooijen/socrates-discord-bot/handlers"
	"github.com/lucianonooijen/socrates-discord-bot/parser"
)

// This function will be called (due to AddHandler) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreateDiogenes(s *discordgo.Session, m *discordgo.MessageCreate) {
	const botPrefix = "!diogenes"

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Notify users to use the botPrefix and not just mention
	for _, mention := range m.Mentions {
		if mention.ID == s.State.User.ID {
			if m.Type == discordgo.MessageTypeReply { // Ignore all messages that are replies that mention @Socrates
				return
			}
			summonInfo := fmt.Sprintf("To summon me, please use %s (instead of mentioning me)", botPrefix)
			s.ChannelMessageSendReply(m.ChannelID, summonInfo, m.MessageReference)
		}
	}

	// Only listen to messages starting with the correct prefix
	contents := strings.ToLower(m.Content)

	if !strings.HasPrefix(contents, botPrefix) {
		fmt.Printf("message '%s' does not have prefix %s\n", contents, botPrefix)
		handlers.KeywordsDiogenesEmoji(s, m, contents)
		handlers.KeywordsDiogenesText(s, m, contents)
		return
	}
	fmt.Printf("Found a message starting with '%s': '%s'\n", botPrefix, contents)

	// Check if the channel is NSFW marked
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		handlers.ReportErrorMessage(s, m.ChannelID, err)
	}
	if !channel.NSFW {
		notNswf := "For obvious reasons, I will only respond to you in channels marked as NSFW"
		s.ChannelMessageSendReply(m.ChannelID, notNswf, m.MessageReference)
		return
	}

	// Get the command given and run the correct handler
	commandString := strings.TrimSpace(strings.TrimPrefix(contents, botPrefix))
	botCommand, args := parser.ParseRequest(commandString)
	fmt.Printf("detected bot command '%s' with args %#v\n", botCommand, args)

	switch botCommand {
	case "speak":
		handlers.SpeakNsfw(s, m)
	case "redpill":
		handlers.Redpill(s, m)
	case "source":
		handlers.Source(s, m)
	case "help":
		handlers.HelpDiogenes(s, m)
	case "supersecretreplycommand":
		handlers.AnonymousReply(s, args)
	default:
		handlers.HelpDiogenes(s, m)
	}
}
