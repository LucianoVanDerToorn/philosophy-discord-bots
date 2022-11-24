package commands

import (
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/handlers"
	"github.com/bwmarrin/discordgo"
)

func moderatorCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	subcommand := data.Options[0].Name
	switch subcommand {
	case commandApprove:
		member := data.Options[0].Options[0].Value
		if member == nil {
			sendInteractionError(s, i, "Member value not set")
			return
		}
		user := member.(string)
		handlers.ApproveMember(s, i, user)

		//InteractionMessageResponse(s, i, fmt.Sprintf("Member %s approved, data %#v", memberToApprove, data))
	default:
		sendInteractionError(s, i, "Something went wrong, could not parse command")
	}
}
