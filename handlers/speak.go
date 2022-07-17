package handlers

import (
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/speak"
	"github.com/bwmarrin/discordgo"
)

func SpeakSfw(s *discordgo.Session, m *discordgo.MessageCreate) {
	quote := speak.RandomQuote(false)
	s.ChannelMessageSend(m.ChannelID, quote)
}

func SpeakNsfw(s *discordgo.Session, m *discordgo.MessageCreate) {
	quote := speak.RandomQuote(true)
	s.ChannelMessageSend(m.ChannelID, quote)
}
