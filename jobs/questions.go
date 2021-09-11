package jobs

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/lucianonooijen/socrates-discord-bot/handlers"
	"github.com/lucianonooijen/socrates-discord-bot/internal/questions"
	"github.com/robfig/cron/v3"
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

	_ = s.MessageReactionAdd(questionChannelId, msg.ID, "yesvote")
	_ = s.MessageReactionAdd(questionChannelId, msg.ID, "novote")
}
