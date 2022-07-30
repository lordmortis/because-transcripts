package main

import (
	"database/sql"
	"gopkg.in/errgo.v2/errors"
	"os"
	"runtime"

	"github.com/jessevdk/go-flags"

	"BecauseLanguageBot/config"
)

type Options struct {
	ConfigFile string `long:"configFile" description:"path to config.json file" default:"../config.yaml"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)
var configFile *config.Config
var dbCon *sql.DB

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	parser.CommandHandler = func(command flags.Commander, args []string) error {
		var err error

		configFile, err = config.Load(&options.ConfigFile)
		if err != nil {
			return errors.Because(err, nil, "Unable to parse config file")
		}

		if command == nil {
			return errors.New("No command specified")
		}

		return command.Execute(args)
	}

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
