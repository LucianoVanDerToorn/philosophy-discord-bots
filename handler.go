package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
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
	botCommand := func() string { // TODO: support arguments
		parts := strings.Split(m.Content, " ")
		if len(parts) <= 1 {
			return ""
		}
		return parts[1]
	}()
	fmt.Printf("detected bot command '%s'\n", botCommand)
	switch botCommand {
	case "groupinfo":
		botCommandGroupinfo(s, m) // TODO: Support !socrates groupinfo [name]
	case "speak":
		botSpeak(s, m)
	case "help":
		botCommandHelp(s, m)
	default:
		botCommandHelp(s, m)
	}
}
