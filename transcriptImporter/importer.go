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
	groupSpeakerRegex       *regexp.Regexp
	spokenLineRegex         *regexp.Regexp
	paralinguisticLineRegex *regexp.Regexp
)

func init() {
	groupSpeakerRegex = regexp.MustCompile(`[Aa][Nn][Dd]`)
	spokenLineRegex = regexp.MustCompile(`^(([A-Za-z]*\s?[A-Za-z]*\s?[A-Za-z]*\s?):)(.*)$`)
	paralinguisticLineRegex = regexp.MustCompile(`^\[.*]\s$`)
}

func doImport(filePath string, source *datasource.DataSource) error {
	ctx := context.Background()

	fileDesc, err := os.Open(filePath)
	defer fileDesc.Close()
	if err != nil {
		return errors.Because(err, nil, "could not import file")
	}

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
	lineIndex := 0
	sequenceNo := 0

	groupUtterances := make([]*datasource.Utterance, 0)
	currentSpeakers := make([]*datasource.Speaker, 0, 1)
	currentSpeakerAll := false

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			lineIndex++
			continue
		}

		turn := episode.NewTurn()
		turn.SequenceNo = sequenceNo
		updated, err := turn.Update(ctx)
		if err != nil {
			errorString := fmt.Sprintf("could not create turn for %s:%d", filePath, lineIndex)
			return errors.Because(err, nil, errorString)
		}

		if !updated {
			return errors.New(fmt.Sprintf("could not create turn for %s:%d", filePath, lineIndex))
		}

		sequenceNo += 10

		var utterances []*datasource.Utterance

		if paralinguisticLineRegex.MatchString(line) {
			utterance := handleParalinguistic(turn, 0, line)
			updated, err := utterance.Update(ctx)
			if err != nil {
				errorString := fmt.Sprintf("could not create utterance for %s:%d", filePath, lineIndex)
				return errors.Because(err, nil, errorString)
			}
			if !updated {
				return errors.New(fmt.Sprintf("could not create utterance for %s:%d", filePath, lineIndex))
			}

			lineIndex++
			continue
		} else if spokenLineRegex.MatchString(line) {
			currentSpeakers = currentSpeakers[:0]
			matches := spokenLineRegex.FindAllStringSubmatch(line, -1)
			if len(matches) != 1 || len(matches[0]) != 4 {
				return errors.New(fmt.Sprintf("could not parse %s:%d", filePath, lineIndex))
			}

			utterances, err = handleSpokenLine(turn, strings.TrimSpace(matches[0][3]))
			if err != nil {
				return errors.New(fmt.Sprintf("could not create utterances from %s:%d", filePath, lineIndex))
			}

			transcriptName := matches[0][2]
			if strings.EqualFold(transcriptName, "all") {
				groupUtterances = append(groupUtterances, utterances...)
				currentSpeakerAll = true
			} else {
				currentSpeakerAll = false
				speakerStrings := []string{transcriptName}
				if groupSpeakerRegex.MatchString(transcriptName) {
					speakerMatches := groupSpeakerRegex.FindAllStringIndex(transcriptName, -1)
					speakerStrings = []string{strings.TrimSpace(transcriptName[0:speakerMatches[0][0]])}
					speakerStrings = append(speakerStrings, strings.TrimSpace(transcriptName[speakerMatches[0][1]+1:]))
				}

				for len(speakerStrings) > 0 {
					transcriptName = speakerStrings[0]
					speakerStrings = speakerStrings[1:]
					currentSpeaker, err := source.SpeakerWithTranscriptName(ctx, transcriptName)
					if err != nil {
						errorString := fmt.Sprintf("could not find speaker %s for line %d from file '%s'", transcriptName, lineIndex, filePath)
						return errors.Because(err, nil, errorString)
					}

					if currentSpeaker == nil {
						currentSpeaker = source.NewSpeaker()
						currentSpeaker.TranscriptName = transcriptName
						currentSpeaker.Name = transcriptName
						_, err := currentSpeaker.Update(ctx)
						if err != nil {
							errorString := fmt.Sprintf("could not create speaker %s for line %d from file '%s'", transcriptName, lineIndex, filePath)
							return errors.Because(err, nil, errorString)
						}
					}
					currentSpeakers = append(currentSpeakers, currentSpeaker)
				}

			}
		} else {
			utterances, err = handleSpokenLine(turn, line)
		}

		for _, utterance := range utterances {
			if !currentSpeakerAll && len(currentSpeakers) > 0 {
				utterance.Speakers = currentSpeakers
			} else {
				utterance.Speakers = []*datasource.Speaker{}
			}

			success, err := utterance.Update(ctx)
			if err != nil {
				errorString := fmt.Sprintf("could not add utterance for line %d from file '%s'", lineIndex, filePath)
				return errors.Because(err, nil, errorString)
			}

			if !success {
				errorString := fmt.Sprintf("could not add utterance for line %d from file '%s'", lineIndex, filePath)
				return errors.New(errorString)
			}
		}

		lineIndex++
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
