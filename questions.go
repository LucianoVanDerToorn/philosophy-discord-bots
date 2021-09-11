package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/lucianonooijen/socrates-discord-bot/handlers"
	"github.com/lucianonooijen/socrates-discord-bot/internal/questions"
	"github.com/robfig/cron/v3"
)

// TODO: Add `jobs` package

const questionChannelId = "878925957602369566"

func addBotActionCron(s *discordgo.Session) {
	c := cron.New()
	_, err := c.AddFunc("0 12 * * *", func() {
		botActionQuestion(s)
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added cron job for ")
	}
	c.Start()
}

func botActionQuestion(s *discordgo.Session) {
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

	_ = s.MessageReactionAdd(questionChannelId, msg.ID, "yesvote")
	_ = s.MessageReactionAdd(questionChannelId, msg.ID, "novote")
}
