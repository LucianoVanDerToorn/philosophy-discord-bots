package handlers

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

const introductionsChannelId = "895401321514016830"
const introductionMessageTemplate = `Welcome, %s! My name is Socrates, the hand-crafted bot for this Philosophy Bookclub server, nice to meet you.

I would like to ask you a few questions. After that, a human moderator will give you access to the server, then you can pick your roles in the <#784552536044339201> channel.

An overview of all reading groups for the upcoming week can be found in <#883007166859083796>.

1. What are your name, pronouns and age?
2. How did you get interested in philosophy?
3. What areas of  philosophy and which thinkers interest you?
4. What kinds of people do you want to have discussions with?
5. Do you agree with the <#784544880591896599> and in particular, the principle of charity?

You can answer these questions in <#784549937345593415>`

func MemberJoins(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	time.Sleep(2 * time.Second) // Wait two seconds to grasp the user's attention

	mention := m.Mention()
	message := fmt.Sprintf(introductionMessageTemplate, mention)

	if _, err := s.ChannelMessageSend(introductionsChannelId, message); err != nil {
		ReportErrorMessage(s, logChannelId, fmt.Errorf("error sending message in MemberJoins handler: %s", err))
	}
}
