package jobs

import (
	"fmt"
	"strings"
	"time"

	"github.com/LucianoVanDerToorn/philosophy-discord-bots/handlers"
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/groupinfo"
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

// Note: times are the times for notification
// calculate this using meeting time minus one hour
// use GMT times.
func AddNotificationCron(s *discordgo.Session) {
	c := cron.New()
	addNotificationCronWeekly(c, s, "30 19 * * 1", "784556092814393394", "arendt", "784565934584889405")
	addNotificationCronWeekly(c, s, "0 18 * * 3", "874324504401301564", "stirner", "874508087598383154")
	//addNotificationCronWeekly(c, s, "0 14 * * 5", "827469681689755688", "nietzsche", "827528609081458718")
	//addNotificationCronWeekly(c, s, "0 15 * * 0", "859164046922481705", "stoicism", "858964070749175808")
	c.Start()
}

func addNotificationCronWeekly(c *cron.Cron, s *discordgo.Session, crontimes string, channelId string, channel string, roleId string) {
	_, err := c.AddFunc(crontimes, notifyJob(s, channelId, channel, roleId))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added notification job for channel %s\n", channel)
}

func shouldSkipNotification(now time.Time, channel string) bool {
	_ = now.Day()
	return false
}

func notifyJob(s *discordgo.Session, channelId string, channel string, roleId string) func() {
	return func() {
		// Do not send notifiction on a few national holidays
		now := time.Now()
		if shouldSkipNotification(now, channel) {
			skipMessage := fmt.Sprintf("There is no %s meeting today.\n(Moderators, feel free to correct me if I'm wrong, because I know nothing)", strings.Title(channel))
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

			if cd.MeetingOn == groupinfo.MeetingOnUnsure {
				_, err := s.ChannelMessageSend(channelId, "The meeting host will shortly let everyone know where the meeting will be held.")
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
