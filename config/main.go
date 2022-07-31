package config

import (
	"github.com/spf13/viper"
	"gopkg.in/errgo.v2/errors"
)

type DiscordConfig struct {
	ClientID string `mapstructure:"client_id"`
	Token    string
}

type DatabaseConfig struct {
	Path string
}

type TranscriptConfig struct {
	Directory            string
	IncludedContextLines int `mapstructure:"included_context_lines"`
}

type ImporterConfig struct {
	Directory string
	WaitTime  int `mapstructure:"wait_time"`
}

type HttpConfig struct {
	BindAddress string `mapstructure:"bind_address"`
	Port        int
}

type Config struct {
	DiscordConfig    DiscordConfig    `mapstructure:"discord"`
	DatabaseConfig   DatabaseConfig   `mapstructure:"database"`
	ImporterConfig   ImporterConfig   `mapstructure:"importer"`
	TranscriptConfig TranscriptConfig `mapstructure:"transcripts"`
	HttpConfig       HttpConfig       `mapstructure:"http"`
	DevelopmentMode  bool             `mapstructure:"development"`
}

func Load(filename *string) (*Config, error) {
	if filename != nil && len(*filename) > 0 {
		viper.SetConfigFile(*filename)
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
