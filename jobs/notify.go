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
	addNotificationCronWeekly(c, s, "30 18 * * 3", "874324504401301564", "stirner", "874508087598383154")
	addNotificationCronWeekly(c, s, "0 18 * * 4", "835711337979838464", "epistemology", "850752244793868298")
	addNotificationCronWeekly(c, s, "0 13 * * 5", "827469681689755688", "nietzsche", "827528609081458718")
	addNotificationCronWeekly(c, s, "0 23 * * 5", "896873495340916736", "taoism", "896875467632693248")
	addNotificationCronWeekly(c, s, "0 12 * * 6", "784555854892367902", "plato", "784565469214146581")
	addNotificationCronWeekly(c, s, "0 14 * * 6", "869272998576787516", "camus", "869240946301235252")
	addNotificationCronWeekly(c, s, "30 12 * * 0", "858318098251513866", "marxism", "858963568955490335")
	addNotificationCronWeekly(c, s, "0 15 * * 0", "859164046922481705", "stoicism", "858964070749175808")
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
		message := fmt.Sprintf("<@&%s> the %s meeting starts in 60 minutes", roleId, strings.Title(channel))
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
