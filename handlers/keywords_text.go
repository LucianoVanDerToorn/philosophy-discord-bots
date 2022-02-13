package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func KeywordsSocratesText(s *discordgo.Session, m *discordgo.MessageCreate, contents string) {
	if len(contents) > 100 {
		return
	}

	contentsLower := strings.ToLower(contents)

	keywords := map[string]string{
		"homer":   "Homer? Talking about Homer? You better go wash your mouth with soap, young one.",
		"1.e4 c5": "play 2.d4 and your opponent can just resign the game.",
	}

	for k, reply := range keywords {
		if strings.Contains(contentsLower, k) {
			_, err := s.ChannelMessageSendReply(m.ChannelID, reply, m.Reference())
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func KeywordsDiogenesText(s *discordgo.Session, m *discordgo.MessageCreate, contents string) {
	if len(contents) > 100 {
		return
	}

	contentsLower := strings.ToLower(contents)
	keywords := map[string]string{
		"meow": "Wowwwww, you meow like a cat! That means you are one, right? Shut the fuck up. If you really want to be put on a leash and treated like a domestic animal then that’s called a fetish, not “quirky” or “cute”. What part of you seriously thinks that any part of acting like a feline establishes a reputation of appreciation? Is it your lack of any defining aspect of personality that urges you to resort to shitty representations of cats to create an illusion of meaning in your worthless life? Wearing “cat ears” in the shape of headbands further notes the complete absence of human attribution to your false sense of personality, such as intelligence or charisma in any form or shape. Where do you think this mindset’s gonna lead you? You think you’re funny, random, quirky even? What makes you think that acting like a fucking cat will make a goddamn hyena laugh? I, personally, feel extremely sympathetic towards you as your only escape from the worthless thing you call your existence is to pretend to be an animal. But it’s not a worthy choice to assert this horrifying fact as a dominant trait, mainly because personality traits require an initial personality to lay their foundation on. You’re not worthy of anybody’s time, so go fuck off, “cat-girl”.",
	}

	for k, reply := range keywords {
		if strings.Contains(contentsLower, k) {
			if strings.Contains(contentsLower, k) {
				_, err := s.ChannelMessageSendReply(m.ChannelID, reply, m.Reference())
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
