package main

import (
	"BecauseLanguageBot/httpServer"
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

	searcher, err = transcriptSearcher.Init(configData.TranscriptConfig)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Transcription directory error: %s\n", err))
		os.Exit(1)
	}

	/*	discordInstance, err := discord.Init(configData.DiscordConfig)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Unable to connect to discord: %s\n", err))
			os.Exit(1)
		}*/

	httpInstance, err := httpServer.Init(configData.HttpConfig)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Unable to connect to discord: %s\n", err))
		os.Exit(1)
	}

	httpInstance.SetDevelopmentMode(configData.DevelopmentMode)

	registerTranscriptSearch(searcher)
	registerInfo()
	httpInstance.Start()

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
}
