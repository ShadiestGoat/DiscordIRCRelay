package main

import (
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var s *discordgo.Session

func main() {
	InitLogger()
	logger.Logf("Logger loaded!")
	InitConfig()
	logger.Logf("Config loaded!")
	go InitIRC()

	var err error
	s, err = discordgo.New("Bot " + DISCORD_TOKEN)
	PanicIfErr(err)

	s.Identify.Intents = discordgo.IntentGuilds |
		discordgo.IntentGuildMessages |
		discordgo.IntentMessageContent

	s.AddHandler(OnDiscordMessage)

	defer s.Close()
	defer logger.File.Close()

	s.AddHandler(func(s *discordgo.Session, e *discordgo.Ready) {
		logger.Logf("Discord connected! Logged in as %v! 6/8", e.User.Username)
	})
	
	err = s.Open()
	PanicIfErr(err)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	
	logger.Logf("Bot is ready! You can use Ctrl+C to shut it down gracefully!")
	
	<-stop
	
	IRCConn.Disconnect()

	logger.Logf("Gracefully shutting down...")
}
