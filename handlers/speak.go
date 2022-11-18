package handlers

import (
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/botid"
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/speak"
	"github.com/bwmarrin/discordgo"
)

func SpeakSocrates(s *discordgo.Session, i *discordgo.InteractionCreate) {
	quote := speak.RandomQuote(botid.Socrates)
	InteractionMessageResponse(s, i, quote)
}

func SpeakDiogenes(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channel, err := s.Channel(i.ChannelID)
	if err != nil {
		ReportErrorMessage(s, i.ChannelID, err)
	}

	if !channel.NSFW {
		InteractionMessageResponse(s, i, "My brother in Christ, do you think summoning me in a non-NSFW channel is a good idea?")
	}

	quote := speak.RandomQuote(botid.Diogenes)
	InteractionMessageResponse(s, i, quote)
}

func SpeakBen(s *discordgo.Session, i *discordgo.InteractionCreate) {
	quote := speak.RandomQuote(botid.Ben)
	InteractionMessageResponse(s, i, quote)
}
