package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/lucianonooijen/socrates-discord-bot/internal/speak"
)

func SpeakSfw(s *discordgo.Session, m *discordgo.MessageCreate) {
	quote := speak.RandomQuote(false)
	s.ChannelMessageSend(m.ChannelID, quote)
}

func SpeakNsfw(s *discordgo.Session, m *discordgo.MessageCreate) {
	quote := speak.RandomQuote(true)
	s.ChannelMessageSend(m.ChannelID, quote)
}
