package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"BecauseLanguageBot/config"
)

var configData *config.Config

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	configPath := flag.String("config", "", "JSON Config File")

	flag.Parse()

	configData, err := config.Load(configPath)

	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Unable to parse config file because: %s", err))
		os.Exit(1)
	}

	fmt.Printf("Config is: %#v", configData)
}
