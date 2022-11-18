package main

import (
	_ "embed"
	"fmt"
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/commands"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/LucianoVanDerToorn/philosophy-discord-bots/handlers"
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/jobs"
)

//go:generate ./version.sh

var (
	//go:embed .version
	Version string
)

func initSocrates(token string) *discordgo.Session {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + strings.ReplaceAll(token, "\n", ""))
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		panic(err)
	}

	// Register the slash commands
	if err = commands.AddCommandsAndHandlersSocrates(dg); err != nil {
		fmt.Println("Error adding Socrates slash commands,", err)
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

func initDiogenes(token string) *discordgo.Session {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + strings.ReplaceAll(token, "\n", ""))
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		panic(err)
	}

	// Register the slash commands
	if err = commands.AddCommandsAndHandlersDiogenes(dg); err != nil {
		fmt.Println("Error adding Diogenes slash commands,", err)
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

func initBenjamin(token string) *discordgo.Session {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + strings.ReplaceAll(token, "\n", ""))
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		panic(err)
	}

	// Register the slash commands
	if err = commands.AddCommandsAndHandlersFinegold(dg); err != nil {
		fmt.Println("Error adding Finegold slash commands,", err)
		panic(err)
	}

	// Register the messageCreateBen func as a callback for MessageCreate events.
	dg.AddHandler(messageCreateBen)

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
				Name: "making Levy cry like a grandmaster",
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
	config := LoadConfigMust()

	sgSocrates := initSocrates(config.Socrates)
	sgDiogenes := initDiogenes(config.Diogenes)
	sgBenjamin := initBenjamin(config.Benjamin)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord sessions.
	sgSocrates.Close()
	sgDiogenes.Close()
	sgBenjamin.Close()
}
