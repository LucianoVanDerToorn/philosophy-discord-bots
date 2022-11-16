package speak

import (
	_ "embed"
	"math/rand"
	"strings"

	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/botid"
	"github.com/LucianoVanDerToorn/philosophy-discord-bots/internal/util"
)

//go:embed speak_socrates.txt
var QuotesFileSocrates string

//go:embed speak_diogenes.txt
var QuotesFileDiogenes string

//go:embed speak_ben.txt
var QuotesFileBen string

func LoadQuotes(botId botid.BotId) []string {
	quotesFile := func() string {
		switch botId {
		case botid.Socrates:
			return QuotesFileSocrates
		case botid.Diogenes:
			return QuotesFileDiogenes
		case botid.Ben:
			return QuotesFileBen
		default:
			return ""
		}
	}()

	lines := strings.Split(quotesFile, "---NEW---")
	return util.RemoveEmptyFromStringArray(lines)
}

func RandomQuote(botId botid.BotId) string {
	quotes := LoadQuotes(botId)
	quotesAmount := len(quotes)
	quoteIndex := rand.Intn(quotesAmount)
	return quotes[quoteIndex]
}
