package main

import (
	"BecauseLanguageBot/datasource"
	"BecauseLanguageBot/httpServer"
	"BecauseLanguageBot/transcriptImporter"
	"BecauseLanguageBot/transcriptSearcher"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"BecauseLanguageBot/config"
)

var configData *config.Config
var searcher *transcriptSearcher.Searcher

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	configPath := flag.String("config", "", "JSON Config File")

	flag.Parse()

	configData, err := config.Load(configPath)

	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Unable to parse config file because: %s\n", err))
		os.Exit(1)
	}

	_, err = datasource.Init(configData.DatabaseConfig, configData.DevelopmentMode)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("database init error: %s\n", err))
		os.Exit(1)
	}

	importer, err := transcriptImporter.Init(configData.ImporterConfig)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Could not setup importer: %s\n", err))
		os.Exit(1)
	}

	httpInstance, err := httpServer.Init(configData.HttpConfig)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Unable to connect to discord: %s\n", err))
		os.Exit(1)
	}

	httpInstance.SetDevelopmentMode(configData.DevelopmentMode)
	httpInstance.Start()

	err = importer.Start()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Could not start importer: %s\n", err))
		os.Exit(1)
	}

	/*	discordInstance, err := discord.Init(configData.DiscordConfig)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Unable to connect to discord: %s\n", err))
			os.Exit(1)
		}*/

	registerTranscriptSearch(searcher)
	registerInfo()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	fmt.Println("To add this bot to your server, visit https://discordapp.com/oauth2/authorize?scope=bot&permissions=o&client_id=" + configData.DiscordConfig.ClientID)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	/*	err = discordInstance.Close()
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Unable to close discord connection cleanly: %s\n", err))
			os.Exit(1)
		}*/

	err = httpInstance.Stop()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Unable to shutdown HTTP server cleanly: %s\n", err))
		os.Exit(1)
	}

	err = importer.Stop()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Unable to shutdown importer server cleanly: %s\n", err))
		os.Exit(1)
	}
}
