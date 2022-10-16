package transcriptImporter

import (
	"fmt"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

const checkInterval = 1 * time.Second

type fileInfo struct {
	name           string
	stabilizedTime time.Time
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
		entry, found := importer.watchedEntries[name]
		if !found {
			break
		}

		if entry.stabilizedTime.Before(time.Now()) {
			importer.workQueue <- name
			delete(importer.watchedEntries, entry.name)
		}
	}
}
