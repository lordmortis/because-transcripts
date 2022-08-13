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
	spokenLineRegex         *regexp.Regexp
	paralinguisticLineRegex *regexp.Regexp
)

func init() {
	spokenLineRegex = regexp.MustCompile(`^(([A-Za-z\s]*):)(.*)$`)
	paralinguisticLineRegex = regexp.MustCompile(`^\[.*]$`)
}

func doImport(filePath string, source *datasource.DataSource) error {
	ctx := context.Background()

	fileDesc, err := os.Open(filePath)
	defer fileDesc.Close()
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
	sequenceNo := 0

	groupUtterances := make([]*datasource.Utterance, 0)
	var currentSpeaker *datasource.Speaker
	currentSpeakerAll := false
	currentSpeaker = nil

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			lineindex++
			continue
		}

		turn := episode.NewTurn()
		turn.SequenceNo = sequenceNo
		updated, err := turn.Update(ctx)
		if err != nil {
			errorString := fmt.Sprintf("could not create turn for %s:%d", filePath, lineindex)
			return errors.Because(err, nil, errorString)
		}

		if !updated {
			return errors.New(fmt.Sprintf("could not create turn for %s:%d", filePath, lineindex))
		}

		sequenceNo += 10

		var utterances []*datasource.Utterance

		if paralinguisticLineRegex.MatchString(line) {
			utterance := handleParalinguistic(turn, 0, line)
			updated, err := utterance.Update(ctx)
			if err != nil {
				errorString := fmt.Sprintf("could not create utterance for %s:%d", filePath, lineindex)
				return errors.Because(err, nil, errorString)
			}
			if !updated {
				return errors.New(fmt.Sprintf("could not create utterance for %s:%d", filePath, lineindex))
			}

			lineindex++
			continue
		} else if spokenLineRegex.MatchString(line) {
			matches := spokenLineRegex.FindAllStringSubmatch(line, -1)
			if len(matches) != 1 || len(matches[0]) != 4 {
				return errors.New(fmt.Sprintf("could not parse %s:%d", filePath, lineindex))
			}

			utterances, err = handleSpokenLine(turn, strings.TrimSpace(matches[0][3]))
			if err != nil {
				return errors.New(fmt.Sprintf("could not create utterances from %s:%d", filePath, lineindex))
			}

			transcriptName := matches[0][2]
			if strings.EqualFold(transcriptName, "all") {
				groupUtterances = append(groupUtterances, utterances...)
				currentSpeaker = nil
				currentSpeakerAll = true
			} else {
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
			}
		} else {
			utterances, err = handleSpokenLine(turn, line)
		}

		for _, utterance := range utterances {
			if !currentSpeakerAll && currentSpeaker != nil {
				utterance.Speakers = []*datasource.Speaker{currentSpeaker}
			} else {
				utterance.Speakers = []*datasource.Speaker{}
			}

			success, err := utterance.Update(ctx)
			if err != nil {
				errorString := fmt.Sprintf("could not add utterance for line %d from file '%s'", lineindex, filePath)
				return errors.Because(err, nil, errorString)
			}

			if !success {
				errorString := fmt.Sprintf("could not add utterance for line %d from file '%s'", lineindex, filePath)
				return errors.New(errorString)
			}
		}

		lineindex++
	}

	episodeSpeakers, _, err := episode.Speakers(ctx, -1, -1)
	if err != nil {
		return errors.Because(err, nil, "could not find all speakers")
	}

	for _, utterance := range groupUtterances {
		utterance.Speakers = episodeSpeakers
		_, err := utterance.Update(ctx)
		if err != nil {
			return errors.Because(err, nil, "could not set group speakers")
		}
	}

	fileDesc.Close()
	return nil
}
