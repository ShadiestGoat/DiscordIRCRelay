package main

import "github.com/bwmarrin/discordgo"

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

const EColor = 0x6e6bee
const EGreen = 0x08dd7e
const EError = 0xA51D2A
const EMBED_EMPTY = "\u200b"

func GEmbed(title string, desc string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       title,
		Description: desc,
		Color:       EColor,
	}
}

func GEmbedError(err string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Error!",
		Description: err,
		Color:       EError,
	}
}
