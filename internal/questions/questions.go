package questions

import (
	_ "embed"
	"strings"
	"time"
)

//go:embed questions.txt
var QuestionFile string

func LoadQuestions() map[int]string {
	return ParseQuestionsString(QuestionFile)
}

func ParseQuestionsString(qs string) map[int]string {
	lines := strings.Split(qs, "\n")

	questionMap := make(map[int]string)
	for l, question := range lines {
		dayCount := l + 1
		questionMap[dayCount] = question
	}
	return questionMap
}

func GetQuestionForDate(questions map[int]string, date time.Time) (question string, found bool) {
	dateCount := date.YearDay()
	question, found = questions[dateCount]
	return
}
