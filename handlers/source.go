package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func Source(s *discordgo.Session, i *discordgo.InteractionCreate) {
	message := "You can find my source code here: https://github.com/LucianoVanDerToorn/philosophy-discord-bots"
	InteractionMessageResponse(s, i, message)
}
