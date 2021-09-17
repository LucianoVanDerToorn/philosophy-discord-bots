package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const logChannelId = "888334591549636688"

func SendToLogchannel(s *discordgo.Session, message string) {
	_, err := s.ChannelMessageSend(logChannelId, message)
	if err != nil {
		fmt.Printf("error sending log message: %s", err)
	}
}
