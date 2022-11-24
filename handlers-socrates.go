package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/LucianoVanDerToorn/philosophy-discord-bots/handlers"
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/parser"
)

// This function will be called (due to AddHandler) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreateSocrates(s *discordgo.Session, m *discordgo.MessageCreate) {
	const botPrefix = "!socrates"

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
		handlers.KeywordsSocratesEmoji(s, m, contents)
		handlers.KeywordsSocratesText(s, m, contents)
		return
	}
	fmt.Printf("Found a message starting with '%s': '%s'\n", botPrefix, contents)

	// Get the command given and run the correct handler
	commandString := strings.TrimSpace(strings.TrimPrefix(contents, botPrefix))
	botCommand, args := parser.ParseRequest(commandString)
	fmt.Printf("detected bot command '%s' with args %#v\n", botCommand, args)

	switch botCommand {
	case "reply":
		handlers.AnonymousReply(s, m, args)
	case "help":
		handlers.HelpSocrates(s, m)
	default:
		handlers.HelpSocrates(s, m)
	}
}

func memberJoins(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	handlers.MemberJoins(s, m)
}
