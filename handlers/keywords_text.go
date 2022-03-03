package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const fapResponse = `Please DO NOT announce to the server when you are going to go masturbate. This has been a reoccurring issue, and I'm not sure why some people have such under developed social skills that they think that a server full of mostly male strangers would need to know that. No one is going to be impressed and give you a high five (especially considering where that hand has been). I don't want to add this to the rules, since it would be embarrassing for new users to see that we have a problem with this, but it is going to be enforced as a rule from now on.

If it occurs, you will be warned, then additional occurrences will be dealt with at the discretion of modstaff. Thanks.`

func KeywordsSocratesText(s *discordgo.Session, m *discordgo.MessageCreate, contents string) {
	if len(contents) > 100 {
		return
	}

	contentsLower := strings.ToLower(contents)

	keywords := map[string]string{
		"homer":   "Homer? Talking about Homer? You better go wash your mouth with soap, young one.",
		"1.e4 c5": "play 2.d4 and your opponent can just resign the game.",
		"fap":     fapResponse,
		"jack":    fapResponse,
		"jerk":    fapResponse,
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
