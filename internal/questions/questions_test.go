package questions_test

import (
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/questions"
	"testing"
	"time"
)

func TestGetQuestionForDate(t *testing.T) {
	tests := []struct {
		forDay   time.Time
		question string
	}{
		{forDay: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC), question: "Should secular countries abolish public holidays based on the Christian tradition and adopt a new calendar?"},
		{forDay: time.Date(2023, 1, 13, 12, 0, 0, 0, time.UTC), question: "Is truth subjectivity?"},
		{forDay: time.Date(2023, 1, 14, 12, 0, 0, 0, time.UTC), question: "Can we hold beliefs without evidence?"},
		{forDay: time.Date(2023, 1, 15, 12, 0, 0, 0, time.UTC), question: "Should Ancient Greek and/or Latin be used more as a replacement for English in academic writing?"},
		{forDay: time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC), question: "Nothing is more humiliating than to see idiots succeed in enterprises we have failed in. (Flaubert)"},
		{forDay: time.Date(2023, 3, 1, 12, 0, 0, 1, time.UTC), question: "Nothing is more humiliating than to see idiots succeed in enterprises we have failed in. (Flaubert)"},
		{forDay: time.Date(2023, 5, 4, 12, 0, 1, 42, time.UTC), question: "Of course we can learn even from novels, Nace Novels that is, ut it isn't the same thing as serious readings. [H.G. Wells, Kipps]"},
		{forDay: time.Date(2023, 10, 24, 17, 13, 0, 0, time.UTC), question: "Gender is a kind of imitation for which there is no original. (Judith Butler)"},
		{forDay: time.Date(2024, 10, 24, 17, 13, 0, 0, time.UTC), question: "Is pornography a form of rape?"},
		{forDay: time.Date(2023, 12, 31, 12, 0, 0, 0, time.UTC), question: "It's more important to be nice than to be important."},
		{forDay: time.Date(2024, 12, 31, 12, 0, 0, 0, time.UTC), question: "Does chess strategy translate to the real world?"},
	}

	for _, tt := range tests {
		qs := questions.LoadQuestions()
		got, found := questions.GetQuestionForDate(qs, tt.forDay)
		if !found {
			t.Errorf("did not find question for day %s (%dth day of year)", tt.forDay, tt.forDay.YearDay())
		}
		if got != tt.question {
			t.Errorf("got wrong question for day %s (%dth day of year), got '%s', but expected '%s'", tt.forDay, tt.forDay.YearDay(), got, tt.question)
		}
	}
}
