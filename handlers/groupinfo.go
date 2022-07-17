package handlers

import (
	"fmt"

	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/groupinfo"
	"github.com/bwmarrin/discordgo"
)

func Groupinfo(s *discordgo.Session, m *discordgo.MessageCreate, channel string) {
	c, err := s.Channel(m.ChannelID)
	if err != nil {
		ReportErrorMessage(s, m.ChannelID, err)
	}

	if channel == "" { // Without argument, use the channel name
		channel = c.Name
	}

	cd, ok := groupinfo.ChannelDataLookup[channel]
	if !ok {
		fmt.Printf("channel data not found for channel %s", channel)
		s.ChannelMessageSend(m.ChannelID, "I could not find the meeting data for the given channel: "+channel)
		return
	}

	s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Content: "Here you go!",
		Embed: &discordgo.MessageEmbed{
			Type:        discordgo.EmbedTypeRich,
			Title:       fmt.Sprintf("%s Dropbox link", cd.Name),
			Description: cd.EmbedDescription(),
			URL:         cd.ResourcesLink,
		},
	})
}
