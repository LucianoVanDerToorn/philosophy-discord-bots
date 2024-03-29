package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ReportErrorMessage(s *discordgo.Session, channelId string, err error) {
	const oopsieWhoopsieCopyPasta = "OOPSIE WOOPSIE!! Uwu We make a fucky wucky!! A wittle fucko boingo! The code monkeys at our headquarters are working VEWY HAWD to fix this!"
	fmt.Printf("something went wrong: %s", err)
	_, sendErr := s.ChannelMessageSend(channelId, fmt.Sprintf("%s\n(err: %s)", oopsieWhoopsieCopyPasta, err))
	if sendErr != nil {
		fmt.Printf("error sending error message to Discord: %s", err)
	}
}

func InteractionMessageResponse(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: message},
	})
	if err != nil {
		ReportErrorMessage(s, i.ChannelID, err)
	}
}
