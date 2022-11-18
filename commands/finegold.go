package commands

import (
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/handlers"
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/constants"
	"github.com/bwmarrin/discordgo"
)

func AddCommandsAndHandlersFinegold(s *discordgo.Session) error {
	_, err := s.ApplicationCommandBulkOverwrite(constants.AppIdFinegold, constants.GuildIdLive, finegoldCommands)
	if err != nil {
		return err
	}

	s.AddHandler(finegoldHandler)

	return nil
}

var finegoldCommands = []*discordgo.ApplicationCommand{
	{
		Name:        commandIdFinegold,
		Description: "Interact with our Finegold bot",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        commandSpeak,
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Let Finegold speak his wisdom",
			},
			{
				Name:        commandSource,
				Type:        discordgo.ApplicationCommandOptionSubCommand,
				Description: "Get the source code for the bot",
			},
		},
	},
}

func finegoldHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	switch data.Name {
	case commandIdFinegold:
		if data.Options == nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "You need to select an option to interact with Finegold",
				},
			})
		} else {
			finegoldCommandHandler(s, i)
		}
	default:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Something went wrong",
			},
		})
	}
}

func finegoldCommandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()
	subcommand := data.Options[0].Name
	switch subcommand {
	case commandSpeak:
		handlers.SpeakBen(s, i)
	case commandSource:
		handlers.Source(s, i)
	default:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Something went wrong",
			},
		})
	}
}
