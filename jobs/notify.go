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

// Note: times are the times for notification
// calculate this using meeting time minus one hou
// use GMT times.
func AddNotificationCron(s *discordgo.Session) {
	c := cron.New()
	addNotificationCronWeekly(c, s, "30 19 * * 3", "874324504401301564", "stirner", "874508087598383154")
	addNotificationCronWeekly(c, s, "0 19 * * 4", "835711337979838464", "ontology", "850752244793868298")
	addNotificationCronWeekly(c, s, "0 14 * * 5", "827469681689755688", "nietzsche", "827528609081458718")
	addNotificationCronWeekly(c, s, "30 18 * * 5", "903398560689700894", "marxism", "858963568955490335")
	addNotificationCronWeekly(c, s, "0 23 * * 5", "896873495340916736", "taoism", "896875467632693248")
	addNotificationCronWeekly(c, s, "0 13 * * 6", "784555854892367902", "plato", "784565469214146581")
	addNotificationCronWeekly(c, s, "0 15 * * 6", "869272998576787516", "existentialism", "869240946301235252")
	addNotificationCronWeekly(c, s, "0 17 * * 6", "899264473834078238", "aristotle", "784565293954760734")
	addNotificationCronWeekly(c, s, "0 16 * * 0", "859164046922481705", "stoicism", "858964070749175808")
	addNotificationCronWeekly(c, s, "30 19 * * 0", "903601581302362205", "intro-to-philosophy", "903668792771620864")
	c.Start()
}

func addNotificationCronWeekly(c *cron.Cron, s *discordgo.Session, crontimes string, channelId string, channel string, roleId string) {
	_, err := c.AddFunc(crontimes, notifyJob(s, channelId, channel, roleId))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added notification job for channel %s\n", channel)
}

func shouldSkipNotification(now time.Time) bool {
	day := now.Day()
	return day == 24 || // Christmas Eve
		day == 25 || // Christmas Day
		day == 31 || // New Year's Eve
		day == 1 // New Year's Day
}

func notifyJob(s *discordgo.Session, channelId string, channel string, roleId string) func() {
	return func() {
		// Do not send notifiction on a few national holidays
		now := time.Now()
		if shouldSkipNotification(now) {
			skipMessage := fmt.Sprintf("Because it's a holiday today, the %s meeting will not be held.\n(Moderators, feel free to correct me if I'm wrong, because I know nothing)", strings.Title(channel))
			_, err := s.ChannelMessageSend(channelId, skipMessage)
			if err != nil {
				handlers.ReportErrorMessage(s, channelId, err)
			}
			return
		}

		// Initial notification one hour before
		message := fmt.Sprintf("<@&%s> the %s meeting starts in 60 minutes", roleId, strings.Title(channel))
		_, err := s.ChannelMessageSend(channelId, message)
		if err != nil {
			handlers.ReportErrorMessage(s, channelId, err)
		}

		// Add message with where the meeting will take place
		cd, ok := groupinfo.ChannelDataLookup[channel]
		if ok {
			if cd.MeetingOn == groupinfo.MeetingOnZoom {
				_, err := s.ChannelMessageSend(channelId, fmt.Sprintf("The Zoom link is: %s", cd.ZoomLink))
				if err != nil {
					handlers.ReportErrorMessage(s, channelId, err)
				}
			}

			if cd.MeetingOn == groupinfo.MeetingOnDiscord {
				_, err := s.ChannelMessageSend(channelId, "The meeting will take place on Discord, in the 'Reading Groups' Voice Channel")
				if err != nil {
					handlers.ReportErrorMessage(s, channelId, err)
				}
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
