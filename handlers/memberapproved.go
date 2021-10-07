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

func MemberAssignedMemberRole(s *discordgo.Session, m *discordgo.GuildMemberUpdate) {
	memberRoleAssigned := false
	for _, r := range m.Roles {
		if r == memberRoleId {
			memberRoleAssigned = true
		}
	}
	if !memberRoleAssigned {
		return
	}

	mention := m.Mention()
	message := fmt.Sprintf(approvedMessageTemplate, mention)

	if _, err := s.ChannelMessageSend(introductionsChannelId, message); err != nil {
		ReportErrorMessage(s, logChannelId, fmt.Errorf("error sending message in MemberAssignedMemberRole handler: %s", err))
	}
}
