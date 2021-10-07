package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const memberRoleId = "891331569472651294"
const approvedMessageTemplate = `Hooray, %s! You have been granted the Member role and can now access the server!

You can pick your roles in <#784552536044339201>.

An overview of all reading groups for the upcoming week can be found in <#883007166859083796>.

Feel free to ask any of the moderators for help!`

func shouldSendWelcomeMessage(roles []string) bool {
	shouldSend := false
	// If a person has more than one role already, return,
	// because we only want to run if a user is new and assigned their first role (being the member role)
	if len(roles) > 1 {
		return false
	}

	// For safety to avoid panics, run in a loop
	for _, r := range roles {
		if r == memberRoleId {
			shouldSend = true
		}
	}
	return shouldSend
}

func MemberAssignedMemberRole(s *discordgo.Session, m *discordgo.GuildMemberUpdate) {
	if shouldSend := shouldSendWelcomeMessage(m.Roles); !shouldSend {
		return
	}

	mention := m.Mention()
	message := fmt.Sprintf(approvedMessageTemplate, mention)

	if _, err := s.ChannelMessageSend(introductionsChannelId, message); err != nil {
		ReportErrorMessage(s, logChannelId, fmt.Errorf("error sending message in MemberAssignedMemberRole handler: %s", err))
	}
}
