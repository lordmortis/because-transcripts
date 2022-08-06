package transcriptImporter

import (
	"BecauseLanguageBot/datasource"
	"bufio"
	"context"
	"fmt"
	"gopkg.in/errgo.v2/errors"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

var (
	spokenLineRegex     *regexp.Regexp
	paralinguisticRegex *regexp.Regexp
)

func init() {
	spokenLineRegex = regexp.MustCompile(`^(([A-Z\s]*):)(.*)$`)
	paralinguisticRegex = regexp.MustCompile(`^\[.*\]$`)
}

func doImport(filePath string, source *datasource.DataSource) error {
	ctx := context.Background()

	fileDesc, err := os.Open(filePath)
	if err != nil {
		return errors.Because(err, nil, "could not import file")
	}

	fmt.Printf("Importing %s\n", filePath)

	podcast, err := source.PodcastNamed(ctx, "Because Language")
	if err != nil {
		return errors.Because(err, nil, "could not find postcast name")
	}
	if podcast == nil {
		podcast = source.NewPodcast()
		podcast.Name = "Because Language"
		updated, err := podcast.Update(ctx)
		if err != nil {
			return errors.Because(err, nil, "could not create default podcast")
		}
		if !updated {
			return errors.New("could not create default podcast")
		}
	}

	episode := podcast.NewEpisode()
	nameParts := strings.Split(path.Base(filePath), "-")
	number, err := strconv.ParseInt(strings.TrimSpace(nameParts[0]), 10, 0)
	if err != nil {
		return errors.Because(err, nil, "could not parse episode number")
	} else {
		episode.Number = int(number)
	}

	if len(nameParts) > 1 {
		trimmedName := strings.TrimSpace(nameParts[1])
		if len(trimmedName) > 0 {
			episode.Name = trimmedName
		}
	}

	updated, err := episode.Update(ctx)
	if err != nil {
		return errors.Because(err, nil, "could not create episode")
	}
	if !updated {
		return errors.New("could not create episode")
	}

	scanner := bufio.NewScanner(fileDesc)
	lineindex := 0

	var currentSpeaker *datasource.Speaker
	currentSpeaker = nil

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		if paralinguisticRegex.MatchString(line) {
			fmt.Printf("Paralinguistic: %s\n", line)
		} else if spokenLineRegex.MatchString(line) {
			matches := spokenLineRegex.FindAllStringSubmatch(line, -1)
			if len(matches) != 1 || len(matches[0]) != 4 {
				return errors.New(fmt.Sprintf("could not import speaker on line %d from file '%s'", lineindex, filePath))
			}
			transcriptName := matches[0][2]
			currentSpeaker, err = source.SpeakerWithTranscriptName(ctx, transcriptName)
			if err != nil {
				errorString := fmt.Sprintf("could not find speaker %s for line %d from file '%s'", transcriptName, lineindex, filePath)
				return errors.Because(err, nil, errorString)
			}

			if currentSpeaker == nil {
				currentSpeaker = source.NewSpeaker()
				currentSpeaker.TranscriptName = transcriptName
				currentSpeaker.Name = transcriptName
				_, err := currentSpeaker.Update(ctx)
				if err != nil {
					errorString := fmt.Sprintf("could not create speaker %s for line %d from file '%s'", transcriptName, lineindex, filePath)
					return errors.Because(err, nil, errorString)
				}
			}

			line := matches[0][3]
			_ = line
		} else {
			fmt.Printf("Other line: %s\n", line)
		}

		lineindex++
	}

	defer fileDesc.Close()
	return nil
}
