package speak

import (
	_ "embed"
	"math/rand"
	"strings"

	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/util"
)

//go:embed speak_sfw.txt
var QuotesFileSfw string

//go:embed speak_nsfw.txt
var QuotesFileNsfw string

func LoadQuotes(nsfw bool) []string {
	quotesFile := func() string {
		if nsfw {
			return QuotesFileNsfw
		}
		return QuotesFileSfw
	}()

	lines := strings.Split(quotesFile, "---NEW---")
	return util.RemoveEmptyFromStringArray(lines)
}

func RandomQuote(nswf bool) string {
	quotes := LoadQuotes(nswf)
	quotesAmount := len(quotes)
	quoteIndex := rand.Intn(quotesAmount)
	return quotes[quoteIndex]
}
