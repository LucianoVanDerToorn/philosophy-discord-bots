package handlers

import (
	"fmt"
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/constants"
	"github.com/bwmarrin/discordgo"
)

const approvedMessageTemplate = `Hooray, %s! You have been granted the Member role and can now access the server!

You can pick your roles in <#784552536044339201>.

An overview of all reading groups for the upcoming week can be found in the events tab.

Feel free to ask any of the moderators for help!`

func sendNewMemberWelcomeMessage(s *discordgo.Session, m *discordgo.User) {
	mention := m.Mention()
	message := fmt.Sprintf(approvedMessageTemplate, mention)

	if _, err := s.ChannelMessageSend(introductionsChannelId, message); err != nil {
		ReportErrorMessage(s, logChannelId, fmt.Errorf("error sending message in MemberAssignedMemberRole handler: %s", err))
	}
}

func ApproveMember(s *discordgo.Session, i *discordgo.InteractionCreate, user string) {
	userInfo, err := s.User(user)
	if err != nil {
		InteractionMessageResponse(s, i, fmt.Sprintf("error fetching user %s: %s", user, err))
		return
	}

	username := userInfo.Username

	err = s.GuildMemberRoleAdd(constants.GuildIdLive, user, constants.RoleIdMember)
	if err != nil {
		InteractionMessageResponse(s, i, fmt.Sprintf("error approving member %s (%s): %s", user, username, err))
		return
	}

	sendNewMemberWelcomeMessage(s, userInfo)

	InteractionMessageResponse(s, i, fmt.Sprintf("Approved user %s", username))
}
