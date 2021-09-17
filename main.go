package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/lucianonooijen/socrates-discord-bot/handlers"
	"github.com/lucianonooijen/socrates-discord-bot/jobs"
)

//go:generate ./version.sh

var (
	//go:embed .discordkey
	Token string

	//go:embed .version
	Version string
)

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + strings.ReplaceAll(Token, "\n", ""))
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Register custom functions based on a cronjob
	jobs.AddQuestionCron(dg)
	jobs.AddNotificationCron(dg)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
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
		fmt.Println("Error setting status,", err)
		return
	}

	// Send a message that the bot has come online
	handlers.SendToLogchannel(dg, fmt.Sprintf("I just got online with my last commit hash being %s", Version))

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
