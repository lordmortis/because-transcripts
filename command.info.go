package main

import (
	"BecauseLanguageBot/discord"
	"github.com/bwmarrin/discordgo"
)

func registerInfo() {
	discord.RegisterCommand("info", "information about this bot", handleInfo)
}

func handleInfo(session *discordgo.Session, message *discordgo.Message, parameters string) {
	channel, err := session.UserChannelCreate(message.Author.ID)
	if err != nil {
		discord.LogPMCreateError(message.Author)
		return
	}

	discord.LogChatCommand(message.Author, "Info Request")
	session.ChannelMessageSend(channel.ID, "This is the Because Language Transcript search bot.\nMy code lives at https://github.com/lordmortis/because-transcripts")
}
