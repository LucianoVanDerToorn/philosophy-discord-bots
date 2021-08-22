package questions

import (
	_ "embed"
	"strings"
	"time"
)

//go:embed questions.txt
var QuestionFile string

func LoadQuestions() map[string]string {
	return ParseQuestionsString(QuestionFile)
}

func ParseQuestionsString(qs string) map[string]string {
	lines := strings.Split(qs, "\n")

	questionMap := make(map[string]string)
	for _, l := range lines {
		parts := strings.Split(l, "|")
		if len(parts) < 2 {
			continue
		}
		dateString := parts[0]
		question := parts[1]
		questionMap[dateString] = question
	}
	return questionMap
}

func GetQuestionForDate(questions map[string]string, date time.Time) (question string, found bool) {
	dateString := date.Format("2006-01-02")
	question, found = questions[dateString]
	return
}
