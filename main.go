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
	//go:embed socrates.discordkey
	TokenSocrates string

	//go:embed diogenes.discordkey
	TokenDiogenes string

	//go:embed .version
	Version string
)

func initSocrates() *discordgo.Session {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + strings.ReplaceAll(TokenSocrates, "\n", ""))
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		panic(err)
	}

	// Register the messageCreateSocrates func as a callback for MessageCreate events.
	dg.AddHandler(messageCreateSocrates)

	// Register the memberJoins func as a callback for when a new member joins the server.
	dg.AddHandler(memberJoins)

	// Register the memberGetsMemberRole func as a callback when a member gets a role
	dg.AddHandler(memberGetsMemberRole)

	// Register custom functions based on a cronjob
	jobs.AddQuestionCron(dg)
	jobs.AddNotificationCron(dg)

	// Set the intents
	dg.Identify.Intents =
		discordgo.IntentsAllWithoutPrivileged |
			discordgo.IntentsGuildMembers

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		panic(err)
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
		panic(err)
	}

	// Send a message that the bot has come online
	handlers.SendToLogchannel(dg, fmt.Sprintf("I just got online with my last commit hash being %s", Version))
	return dg
}

func initDiogenes() *discordgo.Session {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + strings.ReplaceAll(TokenDiogenes, "\n", ""))
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		panic(err)
	}

	// Register the messageCreateDiogenes func as a callback for MessageCreate events.
	dg.AddHandler(messageCreateDiogenes)

	// Set the intents
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		panic(err)
	}

	// Set status and activity
	err = dg.UpdateStatusComplex(discordgo.UpdateStatusData{
		IdleSince: nil,
		Activities: []*discordgo.Activity{
			{
				Name: "searching for a man",
				Type: 5,
			},
		},
		AFK:    false,
		Status: "online",
	})
	if err != nil {
		fmt.Println("Error setting status,", err)
		panic(err)
	}

	// Send a message that the bot has come online
	handlers.SendToLogchannel(dg, fmt.Sprintf("I just got online with my last commit hash being %s", Version))
	return dg
}

func main() {
	sgSocrates := initSocrates()
	sgDiogenes := initDiogenes()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord sessions.
	sgSocrates.Close()
	sgDiogenes.Close()
}
