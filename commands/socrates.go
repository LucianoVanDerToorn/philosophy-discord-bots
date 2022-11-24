package commands

import (
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/handlers"
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/constants"
	"github.com/bwmarrin/discordgo"
)

func AddCommandsAndHandlersSocrates(s *discordgo.Session) error {
	_, err := s.ApplicationCommandBulkOverwrite(constants.AppIdSocrates, constants.GuildIdLive, socratesCommands)
	if err != nil {
		return err
	}

	s.AddHandler(socratesHandler)

	return nil
}

var socratesCommands = []*discordgo.ApplicationCommand{
	{
		Name:        commandIdSocrates,
		Description: "Interact with our Socrates bot",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        commandSpeak,
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Let Socrates speak his wisdom",
			},
			{
				Name:        commandSource,
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Get the source code for the bot",
			},
		},
	},

	// Moderator commands
	{
		Name:        commandIdMod,
		Description: "Moderator commands",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        commandApprove,
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Approve a member",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionUser,
						Name:        "user",
						Description: "User to approve",
						Required:    true,
					},
				},
			},
		},
	},
}

func socratesHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	switch data.Name {
	case commandIdSocrates:
		if data.Options == nil {
			sendInteractionError(s, i, "You need to select an option to interact with Socrates")
		} else {
			socratesCommandHandler(s, i)
		}
	case commandIdMod:
		if data.Options == nil {
			sendInteractionError(s, i, "You need to select an option to interact with Moderator commands")
		} else {
			moderatorCommandHandler(s, i)
		}
	default:
		sendInteractionError(s, i, "Something went wrong; commandId not recognized")
	}
}

func socratesCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	subcommand := data.Options[0].Name
	switch subcommand {
	case commandSpeak:
		handlers.SpeakSocrates(s, i)
	case commandSource:
		handlers.Source(s, i)
	default:
		sendInteractionError(s, i, "Something went wrong")
	}
}
