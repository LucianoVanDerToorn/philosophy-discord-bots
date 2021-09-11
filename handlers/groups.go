package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/lucianonooijen/socrates-discord-bot/internal/groupinfo"
)

func Groups(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "**Available groups:**\n"
	for name, data := range groupinfo.ChannelDataLookup {
		groupLine := fmt.Sprintf("* #%s (reading: %s)\n", name, data.ReadingWhat)
		message += groupLine
	}
	s.ChannelMessageSend(m.ChannelID, message)
}
