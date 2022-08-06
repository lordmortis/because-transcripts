package transcriptImporter

import (
	"BecauseLanguageBot/datasource"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/alphadose/zenq/v2"
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

func initialImportRoutine(dataSource *datasource.DataSource, filePath string, wg *sync.WaitGroup, errorQueue *zenq.ZenQ[error]) {
	defer wg.Done()
	err := doImport(filePath, dataSource)
	if err == nil {
		return
	}
	errorString := fmt.Sprintf("could not import '%s'", filePath)
	if errorQueue.IsClosed() {
		return
	}
	errorQueue.Write(errors.Because(err, nil, errorString))
}

func (importer *Importer) Start() error {
	var err error

	importDirEntries, err := os.ReadDir(importer.directory)
	if err != nil {
		errorString := fmt.Sprintf("Could open '%s'", importer.directory)
		return errors.Because(nil, err, errorString)
	}

	var wg sync.WaitGroup
	errorQueue := zenq.New[error](uint32(len(importDirEntries)))
	for _, file := range importDirEntries {
		wg.Add(1)
		filePath := fmt.Sprintf("%s%c%s", importer.directory, os.PathSeparator, file.Name())
		go initialImportRoutine(importer.datasource, filePath, &wg, errorQueue)
	}

	wg.Wait()

	for {
		data, queueOpen := errorQueue.Read()
		if !queueOpen {
			break
		}
		if data == nil {
			continue
		}
		errorQueue.Close()
		return data
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
