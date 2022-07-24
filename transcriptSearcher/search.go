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
	preContext  []string
	matching    []string
	postContext []string
	Pre         []Line
	Matching    []Line
	Post        []Line
}

type Line struct {
	IsInterjection bool
	Speaker        string
	Text           string
}

type searchContext struct {
	searchConfig *Searcher
	regex        *regexp.Regexp
}

var (
	spokenLineRegex *regexp.Regexp
	annotationRegex *regexp.Regexp
)

func init() {
	spokenLineRegex = regexp.MustCompile(`^(.*):(.*)$`)
	annotationRegex = regexp.MustCompile(`^\[.*\]$`)
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
		preContext:  make([]string, 0, config.linesOfContext),
		matching:    make([]string, 0, 1),
		postContext: make([]string, 0, config.linesOfContext),
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
			lastResult.Pre = processIntoLines(lastResult.preContext)
			lastResult.Matching = processIntoLines(lastResult.matching)
			lastResult.Post = processIntoLines(lastResult.postContext)

			lastResult = nil
		}

		if context.regex.MatchString(line) {
			currentResult.matching = append(currentResult.matching, line)
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

	if lastResult != nil {
		lastResult.Pre = processIntoLines(lastResult.preContext)
		lastResult.Matching = processIntoLines(lastResult.matching)
		lastResult.Post = processIntoLines(lastResult.postContext)
	}

	if len(currentResult.matching) > 0 {
		currentResult.Pre = processIntoLines(lastResult.preContext)
		currentResult.Matching = processIntoLines(lastResult.matching)
		currentResult.Post = processIntoLines(lastResult.postContext)
		results = append(results, currentResult)
	}

	return results
}

func processIntoLines(lineStrings []string) []Line {
	lines := make([]Line, len(lineStrings))
	for index, lineString := range lineStrings {
		matches := spokenLineRegex.FindAllStringSubmatch(lineString, -1)
		if len(matches) == 0 {
			lines[index].Text = lineString
			continue
		}
		lines[index].Speaker = matches[0][1]
		lines[index].Text = matches[0][2]
	}
	return lines
}

func (result Result) String() string {
	stringVal := strings.Join(result.preContext, "\n")
	stringVal += "\n"
	stringVal += strings.Join(result.matching, "\n")
	stringVal += "\n"
	stringVal += strings.Join(result.postContext, "\n")
	return stringVal
}
