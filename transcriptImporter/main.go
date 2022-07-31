package transcriptImporter

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"time"

	"gopkg.in/errgo.v2/errors"

	"BecauseLanguageBot/config"
)

const checkInterval = 1 * time.Second

type fileInfo struct {
	name           string
	stabilizedTime time.Time
}

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

	go watch(importer)

	err = importer.watcher.Add(importer.directory)
	if err != nil {
		return errors.Because(err, nil, "Could not start watching import directory")
	}

	return err
}

func watch(importer *Importer) {
	watcher := importer.watcher
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Write == fsnotify.Write {
				entry, found := importer.watchedEntries[event.Name]
				if !found {
					entry = fileInfo{
						name: event.Name,
					}
					go watchEntry(importer, event.Name)
				}

				entry.stabilizedTime = time.Now().Add(importer.waitTime)
				importer.watchedEntries[event.Name] = entry
			}

			if event.Op&fsnotify.Remove == fsnotify.Remove {
				delete(importer.watchedEntries, event.Name)
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			_, _ = os.Stderr.WriteString(fmt.Sprintf("Watch error: %s", err))
		}

	}
}

func watchEntry(importer *Importer, name string) {
	for {
		time.Sleep(checkInterval)
		fmt.Printf("Checking on '%s'\n", name)
		entry, found := importer.watchedEntries[name]
		if !found {
			break
		}

		if entry.stabilizedTime.Before(time.Now()) {
			//TODO: import file
			fmt.Printf("Import file: '%s'\n", entry.name)
			delete(importer.watchedEntries, entry.name)
			break
		}
	}
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
