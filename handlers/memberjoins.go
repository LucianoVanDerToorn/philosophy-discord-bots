package handlers

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

const introductionsChannelId = "895401321514016830"
const introductionMessageTemplate = `Welcome, %s! My name is Socrates, the hand-crafted bot for this Philosophy Bookclub server, nice to meet you.

I would like to ask you a few questions. After that, a human moderator will give you access to the server. When you get access, I will show you where you can pick roles and see our meeting schedule.

1. What are your name, pronouns and age?
2. How did you first get interested in philosophy?
3. What areas of  philosophy and which thinkers do you want to find out more about?
4. What kinds of people do you want to have discussions with?
5. Do you agree with the <#784544880591896599> and in particular, the principle of charity?

You can answer these questions in <#784549937345593415>. Feel free to omit the answers to questions you’re uncomfortable with. That said, providing this information will be helpful to us.`

func MemberJoins(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	time.Sleep(3 * time.Second) // Wait a few seconds to grasp the user's attention

	mention := m.Mention()
	message := fmt.Sprintf(introductionMessageTemplate, mention)

	if _, err := s.ChannelMessageSend(introductionsChannelId, message); err != nil {
		ReportErrorMessage(s, logChannelId, fmt.Errorf("error sending message in MemberJoins handler: %s", err))
	}
}
