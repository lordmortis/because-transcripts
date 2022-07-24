package discord

import (
	"BecauseLanguageBot/config"
	"gopkg.in/errgo.v2/errors"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	session *discordgo.Session
}

func Init(config config.DiscordConfig) (*Discord, error) {
	var err error
	var context Discord
	context.session, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, errors.Because(err, nil, "Could not setup discord session")
	}

	context.session.AddHandler(handleMessage)

	err = context.session.Open()
	if err != nil {
		return nil, errors.Because(err, nil, "Could not setup discord session")
	}

	return &context, nil
}

func (context *Discord) Close() error {
	if context.session == nil {
		return nil
	}

	err := context.session.Close()
	if err != nil {
		return errors.Because(nil, err, "Could not close discord session")
	}

	return nil
}

func discordGetCommand(user *discordgo.User, message *discordgo.MessageCreate) (string, bool) {
	var searchprefixes []string = make([]string, 3)
	var command string
	commandFound := false

	searchprefixes[0] = "<@" + user.ID + ">"
	searchprefixes[1] = "@" + user.Username
	searchprefixes[2] = user.Username

	for _, prefix := range searchprefixes {
		if strings.HasPrefix(message.Content, prefix) {
			command = strings.TrimPrefix(message.Content, prefix)
			commandFound = true
			break
		}
	}

	if commandFound {
		command = strings.Trim(command, " ")
	}

	return command, commandFound
}
