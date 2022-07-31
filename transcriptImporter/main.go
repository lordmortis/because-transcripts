package transcriptImporter

import (
	"fmt"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"gopkg.in/errgo.v2/errors"

	"BecauseLanguageBot/config"
)

type Importer struct {
	directory      string
	waitTime       time.Duration
	watcher        *fsnotify.Watcher
	watchedEntries map[string]fileInfo
}

func Init(config config.ImporterConfig) (*Importer, error) {
	if len(config.Directory) == 0 {
		return nil, errors.New("Transcript directory not specified")
	}

	if config.WaitTime <= 0 {
		return nil, errors.New("wait time is too short or negative. Must be above zero")
	}

	statResult, err := os.Stat(config.Directory)
	if os.IsNotExist(err) {
		err = os.MkdirAll(config.Directory, 0700)
		if err != nil {
			return nil, errors.Because(err, nil, "Could not create import directory")
		}
	} else if !statResult.IsDir() {
		return nil, errors.New(fmt.Sprintf("'%s' is not a directory", config.Directory))
	}

	return &Importer{
		directory:      config.Directory,
		waitTime:       time.Second * time.Duration(config.WaitTime),
		watchedEntries: make(map[string]fileInfo),
	}, nil
}

func (importer *Importer) Start() error {
	var err error
	importer.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return errors.Because(err, nil, "Could not start watching import directory")
	}

	//TODO: check for files already present

	go watch(importer)

	err = importer.watcher.Add(importer.directory)
	if err != nil {
		return errors.Because(err, nil, "Could not start watching import directory")
	}

	return err
}

func (importer *Importer) Stop() error {
	if importer.watcher == nil {
		return nil
	}

	if err := importer.watcher.Close(); err != nil {
		return errors.Because(err, nil, "could not shut down directory watcher")
	}

	return nil
}
