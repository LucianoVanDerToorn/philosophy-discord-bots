package handlers

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

const introductionsChannelId = "895401321514016830"
const introductionMessageTemplate = `Welcome, %s! My name is Socrates, the hand-crafted bot for this Philosophy Bookclub server, nice to meet you.

I would like to ask you a few questions. After that, a human moderator will give you access to the server. When you get access, I will show you where you can pick roles and see our meeting schedule.

1. How old are you?
2. Which <#1044943486564192306> interest you?
3. Do you agree with the <#784544880591896599>? (principle of charity is gone)
4. Optionally, tell us a few things about yourself and how you got interested in philosophy.

You can answer these questions in <#784549937345593415>.`

func MemberJoins(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	time.Sleep(3 * time.Second) // Wait a few seconds to grasp the user's attention

	mention := m.Mention()
	message := fmt.Sprintf(introductionMessageTemplate, mention)

	if _, err := s.ChannelMessageSend(introductionsChannelId, message); err != nil {
		ReportErrorMessage(s, logChannelId, fmt.Errorf("error sending message in MemberJoins handler: %s", err))
	}
}
