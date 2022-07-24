package discord

import (
	"BecauseLanguageBot/config"
	"gopkg.in/errgo.v2/errors"

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
