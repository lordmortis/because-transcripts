package config

import (
	"github.com/spf13/viper"
	"gopkg.in/errgo.v2/errors"
)

type DiscordConfig struct {
	ClientID string `mapstructure:"client_id"`
	Token    string
}

type TranscriptConfig struct {
	Directory            string
	IncludedContextLines int `mapstructure:"included_context_lines"`
}

type Config struct {
	DiscordConfig    DiscordConfig    `mapstructure:"discord"`
	TranscriptConfig TranscriptConfig `mapstructure:"transcripts"`
}

func Load(filename *string) (*Config, error) {
	if filename != nil && len(*filename) > 0 {

	} else {
		//TODO: make this work on windows too.
		viper.AddConfigPath("./")
		viper.AddConfigPath("/etc/because-transcripts")
		viper.AddConfigPath("/usr/local/etc/because-transcripts")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.Because(err, nil, "Could not read config: ")
	}

	config := Config{}

	err = viper.UnmarshalExact(&config)
	if err != nil {
		return nil, errors.Because(err, nil, "Could not parse config")
	}

	return &config, nil
}
