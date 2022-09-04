package main

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"

	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
	"gopkg.in/errgo.v2/errors"

	"BecauseLanguageBot/config"
)

type Options struct {
	DBRootUser string `long:"dbRootUser" description:"root username for the database in config" default:"postgres"`
	DBRootPw   string `long:"dbRootPW" description:"root password for the database in config" default:"rootpassword"`
	ConfigFile string `long:"configFile" description:"path to config.json file" default:"../config.yaml"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)
var configFile *config.Config
var dbCon *sql.DB

func dbConnect(username string, password string, database string) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"user='%s' password='%s' database='%s' host='%s' port=%d sslmode='disable'",
		username,
		password,
		database,
		configFile.DatabaseConfig.Hostname,
		configFile.DatabaseConfig.Port,
	)

	dbCon, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Because(err, nil, "unable to connect to database")
	}

	_, err = dbCon.Exec("SELECT true")
	if err != nil {
		return nil, errors.Because(err, nil, "unable to connect to database")
	}

	return dbCon, nil
}

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
