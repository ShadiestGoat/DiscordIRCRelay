package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	irc "github.com/thoj/go-ircevent"
)

var IRCConn *irc.Connection

func InitIRC() {
	ircConn := irc.IRC(IRC_USER, IRC_USER)
	ircConn.VerboseCallbackHandler = true
	ircConn.Debug = false
	ircConn.UseTLS = false
	ircConn.Password = IRC_PASSWORD

	ircConn.AddCallback("001", func(e *irc.Event) {
		ircConn.Join("#" + RELAY.IRC)
	})

	ircConn.AddCallback("PRIVMSG", func(e *irc.Event) {
		s.WebhookExecute(RELAY.WebhookID, RELAY.WebhookToken, false, &discordgo.WebhookParams{
			Content:         e.Arguments[1],
			Username:        e.Nick,
		})
	})

	err := ircConn.Connect(IRC_HOST)
	PanicIfErr(err)

	IRCConn = ircConn

	logger.Logf("IRC Loaded")

	ircConn.Loop()
}

func OnDiscordMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	if m.WebhookID != "" {
		return
	}
	if m.GuildID == "" {
		return
	}
	if RELAY.Discord != m.ChannelID {
		return
	}

	name := m.Member.Nick
	if name == "" {
		name = m.Author.Username
	}
	nline := strings.Split(m.Content, "\n")
	for _, att := range m.Attachments {
		nline = append(nline, att.URL)
	}
	IRCConn.Nick(strings.ReplaceAll(name, " ", "_"))
	for _, content := range nline {
		IRCConn.Privmsg("#" + RELAY.IRC, content)
	}
}