package main

import "github.com/bwmarrin/discordgo"

func botCommandHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "I'm Socrates, a hand-crafted bot for Bookclub Philosophy. You can ask me anything you like, but I only respond to:\n" +
		"`!socrates help` and I will say what I am saying now\n" +
		"`!socrates groups` and I will list all of the groups on this Discord server for you\n" +
		"`!socrates groupinfo [group]` and I will give you info on that group (group is optional)\n" +
		"`!socrates speak` and I will try to come up with some witty saying\n" +
		"`!socrates source` and I will give you the link to my source code on Github"
	s.ChannelMessageSend(m.ChannelID, message)
}

func botCommandSource(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "You can find my source code here: https://github.com/lucianonooijen/socrates-discord-bot"
	s.ChannelMessageSend(m.ChannelID, message)
}
