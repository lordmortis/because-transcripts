package main

import (
	"BecauseLanguageBot/transcriptSearcher"
	"flag"
	"fmt"
	"os"
	"runtime"

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

	episodeResults, err := searcher.Find("Caffeinated")
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
	}
}
