package jobs

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/lucianonooijen/socrates-discord-bot/handlers"
	"github.com/lucianonooijen/socrates-discord-bot/internal/groupinfo"
	"github.com/robfig/cron/v3"
)

func AddNotificationCron(s *discordgo.Session) {
	c := cron.New()
	addNotificationCronWeekly(c, s, "13 18 * * 6", "886260511354789908", "stoicism", "886262889210593280")
	c.Start()
}

func addNotificationCronWeekly(c *cron.Cron, s *discordgo.Session, crontimes string, channelId string, channel string, roleId string) {
	_, err := c.AddFunc(crontimes, notifyJob(s, channelId, channel, roleId))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added notification job for channel %s\n", channel)
}

func notifyJob(s *discordgo.Session, channelId string, channel string, roleId string) func() {
	return func() {
		// Initial notification one hour before
		message := fmt.Sprintf("<@&%s> the %s meeting is starts in 60 minutes", roleId, strings.Title(channel))
		_, err := s.ChannelMessageSend(channelId, message)
		if err != nil {
			handlers.ReportErrorMessage(s, channelId, err)
		}

		// Add message with the zoom link
		cd, ok := groupinfo.ChannelDataLookup[channel]
		if ok {
			_, err := s.ChannelMessageSend(channelId, fmt.Sprintf("The Zoom link is: %s", cd.ZoomLink))
			if err != nil {
				handlers.ReportErrorMessage(s, channelId, err)
			}
		}

		// Dirty way to report when the group is starting in 5 minutes
		// (dirty is ok here because it's no biggy if the job fails)
		go func() {
			time.Sleep(55 * time.Minute)
			message := fmt.Sprintf("<@&%s> the meeting is starting in 5 minutes!", roleId)
			_, err := s.ChannelMessageSend(channelId, message)
			if err != nil {
				handlers.ReportErrorMessage(s, channelId, err)
			}
		}()
	}
}
