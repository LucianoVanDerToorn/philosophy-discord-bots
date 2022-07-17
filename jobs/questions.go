package jobs

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"

	"github.com/LucianoVanDerToorn/philosophy-discord-bots/handlers"
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/questions"
)

const questionChannelId = "878925957602369566"

func AddQuestionCron(s *discordgo.Session) {
	c := cron.New()
	_, err := c.AddFunc("0 12 * * *", func() {
		question(s)
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added cron job for questions")
	}
	c.Start()
}

func question(s *discordgo.Session) {
	qs := questions.LoadQuestions()
	today := time.Now()
	todaysQuestion, found := questions.GetQuestionForDate(qs, today)

	if !found {
		handlers.ReportErrorMessage(s, questionChannelId, fmt.Errorf("question for today not found"))
		return
	}

	questionMessage := fmt.Sprintf("**Today's question:**\n%s", todaysQuestion)
	msg, err := s.ChannelMessageSend(questionChannelId, questionMessage)
	if err != nil {
		handlers.ReportErrorMessage(s, questionChannelId, err)
	}

	// Add all emojis to the message
	var (
		emojiIdYes      = "yesvote%3A845948680288600074"
		emojiIdNo       = "novote%3A845948680020164609"
		emojiIdQuestion = "%E2%9D%93"
		emojiIdPinch    = "%F0%9F%A4%8F"
		emojiIdAstrix   = "*%EF%B8%8F%E2%83%A3"
	)
	for _, e := range []string{emojiIdYes, emojiIdNo, emojiIdQuestion, emojiIdPinch, emojiIdAstrix} {
		err = s.MessageReactionAdd(questionChannelId, msg.ID, e)
		if err != nil {
			handlers.ReportErrorMessage(s, questionChannelId, err)
		}
	}

	// Generate thread name
	now := time.Now()
	todayStr := now.Format("Monday 02-01-2006")
	threadName := fmt.Sprintf("Discussion of question on \"%s\"", todayStr)

	// Open a public thread
	_, err = s.MessageThreadStart(questionChannelId, msg.ID, threadName, 24*60)
	if err != nil {
		handlers.ReportErrorMessage(s, questionChannelId, err)
	}
}
