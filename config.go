package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type confItem struct {
	Res      *string
	Default  string
	Required bool
}

type DonationRole struct {
	Min    float64
	Max    float64
	XPMult float64
	RoleID string
}

type RelayInfo struct {
	Discord string
	IRC string
	WebhookID string
	WebhookToken string
}

var (
	DISCORD_TOKEN             = ""
	GUILD_ID                  = ""
	IRC_HOST = ""
	IRC_PASSWORD = ""
	IRC_USER = ""
	
	RELAY *RelayInfo
)

func InitConfig() {
	godotenv.Load(".env")
	
	rawRelayChannels := ""

	var confMap = map[string]confItem{
		"DISCORD_TOKEN": {
			Res:      &DISCORD_TOKEN,
			Required: true,
		},
		"GUILD_ID": {
			Res:      &GUILD_ID,
			Required: true,
		},
		"RELAY_CHAN": {
			Res: &rawRelayChannels,
			Required: true,
		},
		"IRC_HOST": {
			Res: &IRC_HOST,
			Required: true,
		},
		"IRC_PASSWORD": {
			Res: &IRC_PASSWORD,
		},
		"IRC_USER": {
			Res: &IRC_USER,
			Required: true,
		},
	}

	for name, opt := range confMap {
		item := os.Getenv(name)

		if item == "" {
			if opt.Required {
				panic(fmt.Sprintf("'%v' is a needed variable, but is not present! Please read the README.md file for more info.", name))
			}
			item = opt.Default
		}

		*opt.Res = item
	}

	spl := strings.Split(rawRelayChannels, ":")
	RELAY = &RelayInfo{
		WebhookID:    spl[0],
		WebhookToken: spl[1],
		Discord:      spl[2],
		IRC:          spl[3],
	}
}
