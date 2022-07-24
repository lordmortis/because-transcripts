package transcriptSearcher

import (
	"bufio"
	"fmt"
	"gopkg.in/errgo.v2/errors"
	"os"
	"regexp"
	"strings"
)

type EpisodeResult struct {
	Name    string
	Results []Result
}

type Result struct {
	preContext    []string
	matchingLines []string
	postContext   []string
}

func (result Result) String() string {
	stringVal := strings.Join(result.preContext, "\n")
	stringVal += "\n"
	stringVal += strings.Join(result.matchingLines, "\n")
	stringVal += "\n"
	stringVal += strings.Join(result.postContext, "\n")
	return stringVal
}

type searchContext struct {
	searchConfig *Searcher
	regex        *regexp.Regexp
}

func (config *Searcher) Find(text string) ([]EpisodeResult, error) {
	var results []EpisodeResult
	var context searchContext
	context.searchConfig = config

	regex, err := regexp.Compile(fmt.Sprintf("^.*((?i)%s).*$", text))
	if err != nil {
		return results, errors.Because(nil, err, "Couldn't compile search string")
	}

	context.regex = regex

	transcriptDirEntries, err := os.ReadDir(config.directory)
	if err != nil {
		errorString := fmt.Sprintf("Could not read transcriptDirEntries")
		return results, errors.Because(nil, err, errorString)
	}

	for _, dirEntry := range transcriptDirEntries {
		//TODO: handle errors better
		episodeResults := context.findInFile(dirEntry)
		if len(episodeResults) == 0 {
			continue
		}

		results = append(results, EpisodeResult{
			Name:    dirEntry.Name(),
			Results: episodeResults,
		})
	}

	return results, nil
}

func (config *Searcher) createEmptyResult() Result {
	return Result{
		preContext:    make([]string, 0, config.linesOfContext),
		matchingLines: make([]string, 0, 1),
		postContext:   make([]string, 0, config.linesOfContext),
	}
}

func (context *searchContext) findInFile(entry os.DirEntry) []Result {
	config := context.searchConfig
	info, err := entry.Info()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Could not read transcript file '%s' info\n", entry.Name()))
		return []Result{}
	}

	if !info.Mode().IsRegular() {
		os.Stderr.WriteString(fmt.Sprintf("Transcript file '%s' is not a regular file\n", entry.Name()))
		return []Result{}
	}

	path := fmt.Sprintf("%s%c%s", config.directory, os.PathSeparator, info.Name())

	fmt.Printf("Starting search in \"%s\"\n", path)
	fileDesc, err := os.Open(path)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Could not read transcript file '%s'\n", entry.Name()))
		return []Result{}
	}
	defer fileDesc.Close()

	results := make([]Result, 0, 5)

	currentResult := config.createEmptyResult()
	var lastResult *Result
	lastResult = nil

	scanner := bufio.NewScanner(fileDesc)
	lineIndex := 0
	contextLineIndex := 0
	for scanner.Scan() {
		lineIndex++

		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		if lastResult != nil && len(lastResult.postContext) < cap(lastResult.postContext) {
			lastResult.postContext = append(lastResult.postContext, line)
		} else if lastResult != nil {
			lastResult = nil
		}

		if context.regex.MatchString(line) {
			currentResult.matchingLines = append(currentResult.matchingLines, line)
			contextLineIndex = 0
			results = append(results, currentResult)
			lastResult = &results[len(results)-1]
			currentResult = config.createEmptyResult()
		} else {
			if contextLineIndex < config.linesOfContext {
				currentResult.preContext = append(currentResult.preContext, line)
			} else {
				currentResult.preContext = append(currentResult.preContext[:0], currentResult.preContext[1:]...)
				currentResult.preContext = append(currentResult.preContext, line)
			}
		}

		contextLineIndex++
	}

	fmt.Printf("ending search in \"%s\"\n", path)
	return results
}
