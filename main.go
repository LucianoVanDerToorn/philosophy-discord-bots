package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Set status and activity
	err = dg.UpdateStatusComplex(discordgo.UpdateStatusData{
		IdleSince: nil,
		Activities: []*discordgo.Activity{
			{
				Name: "debating plebs at the agora",
				Type: 5,
			},
		},
		AFK:    false,
		Status: "online",
	})
	if err != nil {
		fmt.Println("error setting status,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	const botPrefix = "!socrates"

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Notify users to use the botPrefix and not just mention
	for _, mention := range m.Mentions {
		if mention.ID == s.State.User.ID {
			summonInfo := fmt.Sprintf("To summon me, please use %s (instead of mentioning me)", botPrefix)
			s.ChannelMessageSendReply(m.ChannelID, summonInfo, m.MessageReference)
		}
	}

	// Only listen to messages starting with the correct prefix
	contents := m.Content
	if !strings.HasPrefix(contents, botPrefix) {
		fmt.Printf("message '%s' does not have prefix %s\n", contents, botPrefix)
		return
	}
	fmt.Printf("Found a message starting with '%s': '%s'\n", botPrefix, contents)

	// Get the command given and run the correct handler
	botCommand := func()string { // TODO: support arguments
		parts := strings.Split(m.Content, " ")
		if len(parts) <= 1 {
			return ""
		}
		return parts[1]
	}()
	fmt.Printf("detected bot command '%s'\n", botCommand)
	switch botCommand {
	case "groupinfo":
		botCommandGroupinfo(s, m)
	case "help":
		botCommandHelp(s, m)
	default:
		botCommandHelp(s, m)
	}
}

func botCommandHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "I'm Socrates, a hand-crafted bot for Bookclub Philosophy. You can ask me anything you like, but I only respond to `!socrates help` and `!socrates groupinfo` at this moment.")
}

func reportErrorMessage (s *discordgo.Session, channelId string, err error) {
	const oopsieWhoopsieCopyPasta = "OOPSIE WOOPSIE!! Uwu We make a fucky wucky!! A wittle fucko boingo! The code monkeys at our headquarters are working VEWY HAWD to fix this!"
	fmt.Printf("something went wrong: %s", err)
	_, sendErr := s.ChannelMessageSend(channelId, fmt.Sprintf("%s\n(err: %s)", oopsieWhoopsieCopyPasta, err))
	if sendErr != nil {
		fmt.Printf("error sending error message to Discord: %s", err)
	}
}

type ChannelData struct {
	Name string
	Leader string
	MeetingEvery string
	MeetingDay string
	MeetingTimeGmt string
	ReadingWhat string
	DropboxLink string
}
func (cd ChannelData) EmbedDescription() string {
	template := `**Info about the *%s* reading group**
We meet every %s on **%s** at **%s** GMT
We are currently reading **%s**
Press the title to open the Dropbox with the books
The reading group leader is %s for any further questions`
	return fmt.Sprintf(template, cd.Name, cd.MeetingEvery, cd.MeetingDay, cd.MeetingTimeGmt, cd.ReadingWhat, cd.Leader)
}


func botCommandGroupinfo(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.Channel(m.ChannelID)
	if err != nil {
		reportErrorMessage(s, m.ChannelID, err)
	}
	channel := c.Name
	channelDataLookup := map[string]ChannelData {
		"bot-commands": {
			Name:           "Bot Commands",
			Leader:         "Luciano",
			MeetingEvery:   "week",
			MeetingDay:     "Monday",
			MeetingTimeGmt: "10AM",
			ReadingWhat:    "Golang code",
			DropboxLink:    "https://lucianonooijen.com",
		},
	}
	cd, ok := channelDataLookup[channel]
	if !ok {
		fmt.Printf("channel data not found for channel %s", channel)
		s.ChannelMessageSend(m.ChannelID, "I could not find the meeting data for this channel")
		return
	}

	s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Content:         "Here you go!",
		Embed:           &discordgo.MessageEmbed{
			Type:        discordgo.EmbedTypeRich,
			Title:       fmt.Sprintf("%s Dropbox link", cd.Name),
			Description: cd.EmbedDescription(),
			URL: cd.DropboxLink,

		},
	})
}
