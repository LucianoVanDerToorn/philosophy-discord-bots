package handlers

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

type redpill func(s *discordgo.Session, m *discordgo.MessageCreate)

func Redpill(s *discordgo.Session, m *discordgo.MessageCreate) {
	rp := getRedpill()
	rp(s, m)
}

func getRedpill() redpill {
	redpillAmount := len(redpills)
	redpillIndex := rand.Intn(redpillAmount)
	return redpills[redpillIndex]
}

var redpills = []redpill{
	func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			URL:   "http://editions-hache.com/essais/pdf/kaczynski2.pdf",
			Type:  discordgo.EmbedTypeLink,
			Title: "The Industrial Society and It's Future",
		})
	},
	func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			URL:   "https://www.youtube.com/watch?v=mthj2Z7xqvM",
			Type:  discordgo.EmbedTypeVideo,
			Title: "This is The John Birch Society",
		})
	},
	func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			URL:   "https://www.youtube.com/watch?v=YRNKjQg6y-c",
			Type:  discordgo.EmbedTypeVideo,
			Title: "Everything I want to do is illegal",
		})
	},
	func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			URL:   "https://www.youtube.com/watch?v=YRNKjQg6y-c",
			Type:  discordgo.EmbedTypeVideo,
			Title: "Everything I want to do is illegal",
		})
	},
	func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			URL:   "https://www.youtube.com/watch?v=dBnniua6-oM",
			Type:  discordgo.EmbedTypeVideo,
			Title: "Sugar: The Bitter Truth",
		})
	},
}
