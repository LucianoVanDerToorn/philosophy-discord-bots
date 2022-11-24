package permissions

import (
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/constants"
	"github.com/bwmarrin/discordgo"
)

func IsModerator(m *discordgo.Member) bool {
	moderatorRole := constants.RoleIdMod

	for _, r := range m.Roles {
		if r == moderatorRole {
			return true
		}
	}

	return false
}
