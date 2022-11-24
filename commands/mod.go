package commands

import (
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/handlers"
	"github.com/bwmarrin/discordgo"
)

var modCommands = []*discordgo.ApplicationCommand{
	{
		Name:        commandIdMod,
		Description: "Moderator commands",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "test",
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Testing",
			},
		},
	},
}

func moderatorCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	subcommand := data.Options[0].Name
	switch subcommand {
	case "test":
		handlers.SpeakSocrates(s, i)
	default:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Something went wrong, could not parse command",
			},
		})
	}
}
