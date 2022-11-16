package handlers

import "github.com/bwmarrin/discordgo"

func HelpSocrates(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "I'm Socrates, a hand-crafted bot for Bookclub Philosophy. You can ask me anything you like, but I only respond to:\n" +
		"`!socrates help` and I will say what I am saying now\n" +
		"`!socrates groups` and I will list all of the groups on this Discord server for you\n" +
		"`!socrates groupinfo [group]` and I will give you info on that group (group is optional)\n" +
		"`!socrates source` and I will give you the link to my source code on Github"
	s.ChannelMessageSend(m.ChannelID, message)
}

func HelpDiogenes(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "I'm Diogenes, a hand-crafted bot for Bookclub Philosophy, the edgier brother of Socrates. You can talk to me in NSFW channels, and I won't ignore you if you use:\n" +
		"`!diogenes help` and I will repeat myself again like I do now\n" +
		"`!diogenes speak` and I will give you one of my philosophical meditations\n" +
		"`!diogenes redpill` and I will show you the hard truth\n" +
		"`!diogenes source` and I will give you the link to my source code on Github"
	s.ChannelMessageSend(m.ChannelID, message)
}

func HelpFinegold(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "I'm Ben Finegold, a hand-crafted bot for Bookclub Philosophy, the chess bot, my commands:\n" +
		"`!finegold help` and I will repeat myself again like I do now\n" +
		"`!finegold speak` and I will give you one of my philosophical meditations\n" +
		"`!finegold source` and I will give you the link to my source code on Github"
	s.ChannelMessageSend(m.ChannelID, message)
}
