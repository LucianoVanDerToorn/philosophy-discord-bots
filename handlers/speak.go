package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lucianonooijen/socrates-discord-bot/internal/speak"
)

func Speak(s *discordgo.Session, m *discordgo.MessageCreate) {
	quote := speak.RandomQuote()
	s.ChannelMessageSend(m.ChannelID, quote)
}
