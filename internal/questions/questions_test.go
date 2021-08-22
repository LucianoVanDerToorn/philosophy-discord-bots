package questions_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/lucianonooijen/socrates-discord-bot/internal/questions"
)

const testInput = `2021-08-22|Is our world real?
2021-08-23|Do we have free will?
2021-08-24|Are there moral facts?
`

func TestParseQuestionsString(t *testing.T) {
	expected := map[string]string{
		"2021-08-22": "Is our world real?",
		"2021-08-23": "Do we have free will?",
		"2021-08-24": "Are there moral facts?",
	}
	got := questions.ParseQuestionsString(testInput)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("did not get expected result, expected \n%#v\nbut got\n%#v", expected, got)
	}
}

func TestGetQuestionForDate(t *testing.T) {
	qs := questions.ParseQuestionsString(testInput)
	layout := "2006-01-02T15:04:05.000Z"
	str := "2021-08-23T10:49:27.342Z"
	date, err := time.Parse(layout, str)
	if err != nil {
		t.Error(err)
	}

	expected := "Do we have free will?"
	got, found := questions.GetQuestionForDate(qs, date)
	if !found {
		t.Error("did not find question for date")
	}

	if got != expected {
		t.Errorf("got %s but expected %s", got, expected)
	}
}
