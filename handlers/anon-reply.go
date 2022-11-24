package handlers

import (
	"fmt"
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/permissions"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func AnonymousReply(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if !permissions.IsModerator(m.Member) {
		SendToLogchannel(s, fmt.Sprintf("user %s invoked reply command but is not mod", m.Member.Nick))
		return
	}

	if len(args) < 3 {
		SendToLogchannel(s, fmt.Sprintf("to reply, use '[channelid] [messageid] [reply]', %s does not satisfy", args))
		return
	}

	channel := args[0]
	message := args[1]
	reply := strings.Join(args[2:], " ")

	originalMessage, err := s.ChannelMessage(channel, message)
	if err != nil {
		SendToLogchannel(s, fmt.Sprintf("cannot get original message in AnonymousReply: %s", err))
	}

	_, err = s.ChannelMessageSendReply(channel, reply, originalMessage.Reference())
	if err != nil {
		SendToLogchannel(s, fmt.Sprintf("cannot ChannelMessageSendReply in AnonymousReply: %s", err))
	}
}
