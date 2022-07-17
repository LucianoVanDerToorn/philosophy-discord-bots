package handlers

import "github.com/bwmarrin/discordgo"

func Source(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "You can find my source code here: https://github.com/LucianoVanDerToorn/philosophy-discord-bots"
	s.ChannelMessageSend(m.ChannelID, message)
}
