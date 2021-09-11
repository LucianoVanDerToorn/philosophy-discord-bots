package handlers

import "github.com/bwmarrin/discordgo"

func Source(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "You can find my source code here: https://github.com/lucianonooijen/socrates-discord-bot"
	s.ChannelMessageSend(m.ChannelID, message)
}
