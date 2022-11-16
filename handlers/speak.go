package handlers

import (
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/botid"
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/speak"
	"github.com/bwmarrin/discordgo"
)

func SpeakSocrates(s *discordgo.Session, m *discordgo.MessageCreate) {
	quote := speak.RandomQuote(botid.Socrates)
	s.ChannelMessageSend(m.ChannelID, quote)
}

func SpeakDiogenes(s *discordgo.Session, m *discordgo.MessageCreate) {
	quote := speak.RandomQuote(botid.Diogenes)
	s.ChannelMessageSend(m.ChannelID, quote)
}

func SpeakBen(s *discordgo.Session, m *discordgo.MessageCreate) {
	quote := speak.RandomQuote(botid.Ben)
	s.ChannelMessageSend(m.ChannelID, quote)
}
