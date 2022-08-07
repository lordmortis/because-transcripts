package transcriptImporter

import (
	"BecauseLanguageBot/datasource"
	"regexp"
	"strings"
)

var (
	lineParseRegex *regexp.Regexp
)

func init() {
	lineParseRegex = regexp.MustCompile(`(\[.*?])|( [~–—].*?[~–—])|([:?!.…])`)
}

func handleParalinguistic(turn *datasource.Turn, sequenceNo int, rawValue string) *datasource.Utterance {
	utterance := turn.NewUtterance()
	utterance.SequenceNo = sequenceNo
	utterance.IsParalinguistic = true
	utterance.Utterance = rawValue[1 : len(rawValue)-1]
	return utterance
}

func handleUtterance(turn *datasource.Turn, utterances []*datasource.Utterance, sequenceNo int, value string) ([]*datasource.Utterance, int) {
	utteranceString := strings.TrimSpace(value)
	if len(utteranceString) == 0 {
		return utterances, sequenceNo
	}

	utterance := turn.NewUtterance()
	utterance.SequenceNo = sequenceNo
	utterance.IsParalinguistic = false
	utterance.Utterance = utteranceString
	sequenceNo++
	return append(utterances, utterance), sequenceNo
}

func handleSpokenLine(turn *datasource.Turn, rawValue string) ([]*datasource.Utterance, error) {
	matches := lineParseRegex.FindAllStringSubmatchIndex(rawValue, -1)
	if len(matches) == 0 {
		utterances := make([]*datasource.Utterance, 0, 1)
		utterances, _ = handleUtterance(turn, utterances, 0, rawValue)
		return utterances, nil
	}

	sequenceNo := 0
	lastIndex := 0
	utterances := make([]*datasource.Utterance, 0, len(matches))

	for _, match := range matches {
		if match[2] != -1 {
			if match[2] > lastIndex {
				utterances, sequenceNo = handleUtterance(turn, utterances, sequenceNo, rawValue[lastIndex:match[2]-1])
			}
			utterances = append(utterances, handleParalinguistic(turn, sequenceNo, rawValue[match[2]:match[3]]))
			sequenceNo++
		} else if match[4] != -1 {
			if match[4] > lastIndex {
				utterances, sequenceNo = handleUtterance(turn, utterances, sequenceNo, rawValue[lastIndex:match[4]-1])
			}
			utterances, sequenceNo = handleUtterance(turn, utterances, sequenceNo, rawValue[match[4]:match[5]])
		} else if match[6] != -1 {
			utterances, sequenceNo = handleUtterance(turn, utterances, sequenceNo, rawValue[lastIndex:match[7]])
		}
		lastIndex = match[1]
	}

	if lastIndex < len(rawValue) {
		utterances, sequenceNo = handleUtterance(turn, utterances, sequenceNo, rawValue[lastIndex:])
	}

	return utterances, nil
}
