package transcriptSearcher

import (
	"fmt"
	"os"

	"gopkg.in/errgo.v2/errors"

	"BecauseLanguageBot/config"
)

type Searcher struct {
	directory      string
	linesOfContext int
}

func Init(config config.TranscriptConfig) (*Searcher, error) {
	if len(config.Directory) == 0 {
		return nil, errors.New("Transcript directory not specified")
	}

	statResult, err := os.Stat(config.Directory)
	if os.IsNotExist(err) {
		errorString := fmt.Sprintf("Directory %s does not exist", config.Directory)
		return nil, errors.Because(nil, err, errorString)
	}

	if err != nil {
		errorString := fmt.Sprintf("Directory %s could not be accessed", config.Directory)
		return nil, errors.Because(nil, err, errorString)
	}

	if !statResult.IsDir() {
		errorString := fmt.Sprintf("Directory %s is not a directory", config.Directory)
		return nil, errors.Because(nil, err, errorString)
	}

	searcher := Searcher{
		directory:      config.Directory,
		linesOfContext: config.IncludedContextLines,
	}

	transcripts, err := os.ReadDir(config.Directory)
	if err != nil {
		errorString := fmt.Sprintf("Could not read transcripts in '%s'", config.Directory)
		return nil, errors.Because(nil, err, errorString)
	}

	if len(transcripts) == 0 {
		fmt.Printf("Warning: no transcripts in given directory")
	}

	return &searcher, nil
}
