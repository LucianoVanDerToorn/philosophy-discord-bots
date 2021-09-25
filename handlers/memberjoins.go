package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const introductionsChannelId = "865992712511684648"
const introductionMessageTemplate = `Welcome, %s! My name is Socrates, the hand-crafted bot for this Philosophy Bookclub server, nice to meet you.

I would like to ask you a few questions. After that, a human moderator will give you access to the server, then you can pick your roles in the #roles channel.

Feel free to omit the answers to questions you’re uncomfortable with. That said, providing this information will be helpful to us.

1. Name, pronouns, age
2. How did you find us?
3. Education level and area of study (state whether you’re an autodidact, undergrad, graduate and what your major is/was):
4. Specialization (ex: epistemology, philosophy of mind, phenomenology, ethics), if any:
5. Briefly, why are you here?:
6. Which philosophers have you read?
7. Are you committed to the principle of charity? (See Rule 2 listed in #welcome)`

func MemberJoins(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	mention := m.Mention()
	message := fmt.Sprintf(introductionMessageTemplate, mention)

	if _, err := s.ChannelMessageSend(introductionsChannelId, message); err != nil {
		ReportErrorMessage(s, logChannelId, fmt.Errorf("error sending message in MemberJoins handler: %s", err))
	}
}
