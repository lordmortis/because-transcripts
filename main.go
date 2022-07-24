package main

import (
	"BecauseLanguageBot/discord"
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

	discord, err := discord.Init(configData.DiscordConfig)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Unable to connect to discord: %s\n", err))
		os.Exit(1)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	fmt.Println("To add this bot to your server, visit https://discordapp.com/oauth2/authorize?scope=bot&permissions=o&client_id=" + configData.DiscordConfig.ClientID)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	err = discord.Close()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Unable to close discord connection cleanly: %s\n", err))
		os.Exit(1)
	}

	// Cleanly close down the Discord session.

	/*	episodeResults, err := searcher.Find("😟")
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Search Error: %s\n", err))
		}

		totalResults := 0
		for _, episodeResult := range episodeResults {
			totalResults += len(episodeResult.Results)
		}

		if totalResults > 0 {
			fmt.Printf("%d Results\n", totalResults)
		}

		for _, episodeResult := range episodeResults {
			fmt.Printf("Episode: %s:\n", episodeResult.Name)
			for _, result := range episodeResult.Results {
				fmt.Printf("\t Result:\n%s\n", result)
			}
		}*/
}
