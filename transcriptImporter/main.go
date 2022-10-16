package transcriptImporter

import (
	"BecauseLanguageBot/datasource"
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
	datasource     *datasource.DataSource
	workQueue      chan string
}

func Init(config config.ImporterConfig, dataSource *datasource.DataSource) (*Importer, error) {
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
		datasource:     dataSource,
	}, nil
}

func (importer *Importer) ImportWorker() {
	for {
		filePath := <-importer.workQueue
		startTime := time.Now()
		fmt.Printf("Importing '%s\n", filePath)
		err := doImport(filePath, importer.datasource)
		if err != nil {
			fmt.Printf("could not import '%s': %s\n", filePath, err)
			continue
		}

		err = os.Remove(filePath)
		if err != nil {
			fmt.Printf("could not remove file '%s': %s\n", filePath, err)
		}
		endTime := time.Now()
		length := endTime.Sub(startTime)
		fmt.Printf("Imported '%s' in %fs", filePath, length.Seconds())
	}
}

func (importer *Importer) AddToQueue(filename string) {
	importer.workQueue <- fmt.Sprintf("%s%c%s", importer.directory, os.PathSeparator, filename)
}

func (importer *Importer) Start() error {
	var err error

	importDirEntries, err := os.ReadDir(importer.directory)
	if err != nil {
		errorString := fmt.Sprintf("Could open '%s'", importer.directory)
		return errors.Because(nil, err, errorString)
	}

	importer.workQueue = make(chan string)

	go importer.ImportWorker()

	for _, file := range importDirEntries {
		go importer.AddToQueue(file.Name())
	}

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

func (importer *Importer) Stop() error {
	if importer.watcher == nil {
		return nil
	}

	if err := importer.watcher.Close(); err != nil {
		return errors.Because(err, nil, "could not shut down directory watcher")
	}

	return nil
}
