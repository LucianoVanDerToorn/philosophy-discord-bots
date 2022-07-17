package handlers

import (
	"fmt"

	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/groupinfo"
	"github.com/bwmarrin/discordgo"
)

func Groups(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "**Available groups:**\n"
	for name, data := range groupinfo.ChannelDataLookup {
		groupLine := fmt.Sprintf("* #%s (reading: %s)\n", name, data.ReadingWhat)
		message += groupLine
	}
	s.ChannelMessageSend(m.ChannelID, message)
}
