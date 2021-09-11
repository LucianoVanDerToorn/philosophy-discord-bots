package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/lucianonooijen/socrates-discord-bot/handlers"
	"github.com/lucianonooijen/socrates-discord-bot/parser"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	const botPrefix = "!socrates"

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Notify users to use the botPrefix and not just mention
	for _, mention := range m.Mentions {
		if mention.ID == s.State.User.ID {
			summonInfo := fmt.Sprintf("To summon me, please use %s (instead of mentioning me)", botPrefix)
			s.ChannelMessageSendReply(m.ChannelID, summonInfo, m.MessageReference)
		}
	}

	// Only listen to messages starting with the correct prefix
	contents := m.Content
	if !strings.HasPrefix(contents, botPrefix) {
		fmt.Printf("message '%s' does not have prefix %s\n", contents, botPrefix)
		return
	}
	fmt.Printf("Found a message starting with '%s': '%s'\n", botPrefix, contents)

	// Get the command given and run the correct handler
	commandString := strings.TrimSpace(strings.TrimPrefix(contents, botPrefix))
	botCommand, args := parser.ParseRequest(commandString)
	fmt.Printf("detected bot command '%s' with args %#v\n", botCommand, args)

	switch botCommand {
	case "groups":
		handlers.Groups(s, m)
	case "groupinfo":
		if len(args) < 1 {
			handlers.Groupinfo(s, m, "")
		}
		channelName := args[0]
		handlers.Groupinfo(s, m, channelName)
	case "speak":
		handlers.Speak(s, m)
	case "source":
		handlers.Source(s, m)
	case "help":
		handlers.Help(s, m)
	default:
		handlers.Help(s, m)
	}
}
