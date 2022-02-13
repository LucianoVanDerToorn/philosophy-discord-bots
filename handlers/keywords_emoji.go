package handlers

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func KeywordsSocratesEmoji(s *discordgo.Session, m *discordgo.MessageCreate, contents string) {
	if len(contents) > 200 {
		return
	}

	contentsLower := strings.ToLower(contents)
	keywords := map[string]string{
		"putin":       emojiIdYes,
		"socrates":    url.QueryEscape("👀"),
		"pastry":      url.QueryEscape("👎🏽"),
		"pastries":    emojiIdNo,
		"corinth":     url.QueryEscape("👎🏽"),
		"corinthian":  emojiIdNo,
		"homer":       url.QueryEscape("💩"),
		"latin":       url.QueryEscape("🤢"),
		"attic greek": url.QueryEscape("🤙🏽"),
		"meletus":     url.QueryEscape("🙄"),
		"anytus":      url.QueryEscape("🙄"),
		"lycon":       url.QueryEscape("🙄"),
		"sophist":     url.QueryEscape("😠"),
		"alcibiades":  url.QueryEscape("🥵"),
	}

	for k, e := range keywords {
		if strings.Contains(contentsLower, k) {
			err := s.MessageReactionAdd(m.ChannelID, m.ID, e)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func KeywordsDiogenesEmoji(s *discordgo.Session, m *discordgo.MessageCreate, contents string) {
	if len(contents) > 200 {
		return
	}

	contentsLower := strings.ToLower(contents)
	keywords := map[string]string{
		"diogenes":    url.QueryEscape("👀"),
		"honest man":  url.QueryEscape("👀"),
		"latin":       url.QueryEscape("🖕🏽"),
		"attic greek": url.QueryEscape("❤️"),
		"platos man":  url.QueryEscape("🐓"),
		"plato's man": url.QueryEscape("🐓"),
		"dog":         url.QueryEscape("👀️"),
		"corinth":     url.QueryEscape("🍆"),
		"corinthian":  url.QueryEscape("💦"),
	}

	for k, e := range keywords {
		if strings.Contains(contentsLower, k) {
			err := s.MessageReactionAdd(m.ChannelID, m.ID, e)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
