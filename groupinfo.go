package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type ChannelData struct {
	Name           string
	Leader         string
	MeetingEvery   string
	MeetingDay     string
	MeetingTimeGmt string
	ReadingWhat    string
	ResourcesLink  string
	// TODO: Add Zoom link
	// TODO: Add reading next
}

func (cd ChannelData) EmbedDescription() string {
	template := `**Info about the *%s* reading group**
We meet every %s on **%s** at **%s** GMT
We are currently reading **%s**
Press the title to open the link with the books
The reading group leader is %s for any further questions`
	return fmt.Sprintf(template, cd.Name, cd.MeetingEvery, cd.MeetingDay, cd.MeetingTimeGmt, cd.ReadingWhat, cd.Leader)
}

var channelDataLookup = map[string]ChannelData{
	"bot-commands": {
		Name:           "Bot Commands",
		Leader:         "Luciano",
		MeetingEvery:   "week",
		MeetingDay:     "Monday",
		MeetingTimeGmt: "10AM",
		ReadingWhat:    "Golang code",
		ResourcesLink:  "https://lucianonooijen.com",
	},
	"arendt": {
		Name:           "Arendt",
		Leader:         "gintonicelderflower",
		MeetingEvery:   "week",
		MeetingDay:     "Sunday",
		MeetingTimeGmt: "7PM",
		ReadingWhat:    "The Human Condition",
		ResourcesLink:  "https://www.dropbox.com/sh/5nnsx4hyykqq71i/AAAVql4C68PiH_-HbS506VkLa?dl=0",
	},
	"bookworms": {
		Name:           "Bookworms",
		Leader:         "Luciano",
		MeetingEvery:   "month",
		MeetingDay:     "last Saturday of the Month",
		MeetingTimeGmt: "6.30PM",
		ReadingWhat:    "The Shallows (August), Brave New World (September)",
		ResourcesLink:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ", // TODO: Create
	},
	"camus": {
		Name:           "Camus",
		Leader:         "Salman",
		MeetingEvery:   "week",
		MeetingDay:     "Saturdays",
		MeetingTimeGmt: "3PM",
		ReadingWhat:    "The Shallows (August), Brave New World (September)",
		ResourcesLink:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ", // TODO: Create
	},
	"film-discussions": {
		Name:           "Cinephile film discussions",
		Leader:         "Yorgo",
		MeetingEvery:   "month",
		MeetingDay:     "first Saturday of the Month",
		MeetingTimeGmt: "6.30PM",
		ReadingWhat:    "To Live (September)",
		ResourcesLink:  "https://www.youtube.com/watch?v=wNTd0dydfE4",
	},
	"human-nature": {
		Name:           "Cinephile film discussions",
		Leader:         "Yorgo",
		MeetingEvery:   "irregular",
		MeetingDay:     "changing every meeting",
		MeetingTimeGmt: "changing every meeting",
		ReadingWhat:    "changing every meeting",
		ResourcesLink:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ", // TODO: Create
	},
	"marxism": {
		Name:           "Marxism",
		Leader:         "Chris",
		MeetingEvery:   "week",
		MeetingDay:     "Sunday",
		MeetingTimeGmt: "1.30PM",
		ReadingWhat:    "The German Ideology vol. 1",
		ResourcesLink:  "https://www.marxists.org/archive/marx/works/1845/german-ideology/ch01.htm",
	},
	"kripke": {
		Name:           "Kripke (continuing in September)",
		Leader:         "Walker",
		MeetingEvery:   "other week (in odd-weeks)",
		MeetingDay:     "Friday",
		MeetingTimeGmt: "1PM",
		ReadingWhat:    "Naming and Necessity",
		ResourcesLink:  "https://www.dropbox.com/sh/ja5x8nltzhzqz1b/AAC9bWLDcC4tIllJeLWwA8e0a?dl=0",
	},
	"nietzsche": {
		Name:           "Nietzsche",
		Leader:         "Andrew",
		MeetingEvery:   "week",
		MeetingDay:     "Friday",
		MeetingTimeGmt: "6.30PM",
		ReadingWhat:    "On Truth and Lies in an Extra Moral Sense",
		ResourcesLink:  "https://www.dropbox.com/sh/ja5x8nltzhzqz1b/AAC9bWLDcC4tIllJeLWwA8e0a?dl=0",
	},
	"plato": {
		Name:           "Plato",
		Leader:         "Yorgo",
		MeetingEvery:   "week",
		MeetingDay:     "Saturday",
		MeetingTimeGmt: "1PM",
		ReadingWhat:    "The Republic (Book 4/5)",
		ResourcesLink:  "https://www.dropbox.com/sh/ja5x8nltzhzqz1b/AAC9bWLDcC4tIllJeLWwA8e0a?dl=0",
	},
	"sartre": {
		Name:           "Sartre",
		Leader:         "Bob",
		MeetingEvery:   "week",
		MeetingDay:     "Saturday",
		MeetingTimeGmt: "4.30PM",
		ReadingWhat:    "No Exit",
		ResourcesLink:  "https://www.dropbox.com/sh/8x95causpnaka4j/AADZ0HPNKCI3retNNo9uLG0ga?dl=0",
	},
	"stirner": {
		Name:           "Stirner",
		Leader:         "Yorgo",
		MeetingEvery:   "week",
		MeetingDay:     "tbd",
		MeetingTimeGmt: "tbd",
		ReadingWhat:    "tbd",
		ResourcesLink:  "https://www.youtube.com/watch?v=dQw4w9WgXcQ", // TODO: Create
	},
	"stoicism": {
		Name:           "Stoicism",
		Leader:         "Luciano",
		MeetingEvery:   "week",
		MeetingDay:     "Sunday",
		MeetingTimeGmt: "4PM",
		ReadingWhat:    "Epictetus' Enchiridion",
		ResourcesLink:  "https://www.dropbox.com/sh/8hehcb8oda7gc1k/AAC1YE5jwQ7VZK3_mEwkzbCDa?dl=0",
	},
}

func botCommandGroups(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "**Available groups:**\n"
	for name, data := range channelDataLookup {
		groupLine := fmt.Sprintf("* #%s (reading: %s)\n", name, data.ReadingWhat)
		message += groupLine
	}
	s.ChannelMessageSend(m.ChannelID, message)
}

func botCommandGroupinfo(s *discordgo.Session, m *discordgo.MessageCreate, channel string) {
	c, err := s.Channel(m.ChannelID)
	if err != nil {
		reportErrorMessage(s, m.ChannelID, err)
	}

	if channel == "" { // Without argument, use the channel name
		channel = c.Name
	}

	cd, ok := channelDataLookup[channel]
	if !ok {
		fmt.Printf("channel data not found for channel %s", channel)
		s.ChannelMessageSend(m.ChannelID, "I could not find the meeting data for this channel")
		return
	}

	s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Content: "Here you go!",
		Embed: &discordgo.MessageEmbed{
			Type:        discordgo.EmbedTypeRich,
			Title:       fmt.Sprintf("%s Dropbox link", cd.Name),
			Description: cd.EmbedDescription(),
			URL:         cd.ResourcesLink,
		},
	})
}
