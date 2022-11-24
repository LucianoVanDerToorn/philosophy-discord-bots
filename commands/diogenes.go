package commands

import (
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/handlers"
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/constants"
	"github.com/bwmarrin/discordgo"
)

func AddCommandsAndHandlersDiogenes(s *discordgo.Session) error {
	_, err := s.ApplicationCommandBulkOverwrite(constants.AppIdDiogenes, constants.GuildIdLive, diogenesCommands)
	if err != nil {
		return err
	}

	s.AddHandler(diogenesHandler)

	return nil
}

var diogenesCommands = []*discordgo.ApplicationCommand{
	{
		Name:        commandIdDiogenes,
		Description: "Interact with our Diogenes bot",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        commandSpeak,
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Let Diogenes speak his wisdom",
			},
			{
				Name:        commandSource,
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Get the source code for the bot",
			},
		},
	},
}

func diogenesHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	switch data.Name {
	case commandIdDiogenes:
		if data.Options == nil {
			sendInteractionError(s, i, "You need to select an option to interact with Diogenes")
		} else {
			diogenesCommandHandler(s, i)
		}
	default:
		sendInteractionError(s, i, "Something went wrong")
	}
}

func diogenesCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	subcommand := data.Options[0].Name
	switch subcommand {
	case commandSpeak:
		handlers.SpeakDiogenes(s, i)
	case commandSource:
		handlers.Source(s, i)
	default:
		sendInteractionError(s, i, "Something went wrong")
	}
}
