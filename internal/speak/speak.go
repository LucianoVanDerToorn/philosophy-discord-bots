package speak

import (
	_ "embed"
	"math/rand"
	"strings"

	"github.com/lucianonooijen/socrates-discord-bot/internal/util"
)

//go:embed speak.txt
var QuotesFile string

func LoadQuotes() []string {
	lines := strings.Split(QuotesFile, "---NEW---")
	return util.RemoveEmptyFromStringArray(lines)
}

func RandomQuote() string {
	quotes := LoadQuotes()
	quotesAmount := len(quotes)
	quoteIndex := rand.Intn(quotesAmount)
	return quotes[quoteIndex]
}
