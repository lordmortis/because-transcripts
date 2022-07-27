package main

import (
	"BecauseLanguageBot/discord"
	"BecauseLanguageBot/transcriptSearcher"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
)

func registerTranscriptSearch(searcher *transcriptSearcher.Searcher) {
	discord.RegisterCommand("searchTranscripts", "Search transcripts for a given word or phrase", handleTranscriptSearch(searcher))
}

var (
	paralinguisticRegex *regexp.Regexp
)

func init() {
	paralinguisticRegex = regexp.MustCompile(`\[.*\]`)
}

func handleTranscriptSearch(searcher *transcriptSearcher.Searcher) discord.Command {
	return func(session *discordgo.Session, message *discordgo.Message, parameters string) {
		channel, err := session.UserChannelCreate(message.Author.ID)
		if err != nil {
			discord.LogPMCreateError(message.Author)
			return
		}

		discord.LogChatCommand(message.Author, fmt.Sprintf("Transcript Search: '%s'", parameters))

		if len(parameters) == 0 {
			session.ChannelMessageSend(channel.ID, "I need some text to search for")
			return
		}

		session.ChannelMessageSend(channel.ID, fmt.Sprintf("Searching transcripts for '%s'", parameters))

		episodeResults, err := searcher.Find(parameters)
		if err != nil {
			discord.LogChatError(message.Author, "error searching for results", err)
		}

		totalResults := 0
		for _, episodeResult := range episodeResults {
			totalResults += len(episodeResult.Results)
		}

		if totalResults > 0 {
			var msg string
			if totalResults > 20 {
				msg = fmt.Sprintf("%d Results - limiting to the first 20", totalResults)
			} else {
				msg = fmt.Sprintf("%d Results", totalResults)
			}
			session.ChannelMessageSend(channel.ID, msg)
		}

		printedResults := 0

		sendTranscriptLine := func(line transcriptSearcher.Line) {
			msg := ""
			if len(line.Speaker) > 0 {
				msg = fmt.Sprintf("**%s**: %s", cases.Title(language.English).String(line.Speaker), handleTextLine(line.Text))
			} else if line.IsParalinguistic {
				msg = fmt.Sprintf("_%s_", line.Text)
			} else {
				msg = handleTextLine(line.Text)
			}

			session.ChannelMessageSend(channel.ID, msg)
		}

		for _, episodeResult := range episodeResults {
			msg := fmt.Sprintf("Episode: _%s_:", episodeResult.Name)
			session.ChannelMessageSend(channel.ID, msg)

			for index, result := range episodeResult.Results {
				session.ChannelMessageSend(channel.ID, fmt.Sprintf("Result %d", index+1))
				for _, prefix := range result.Pre {
					sendTranscriptLine(prefix)
				}
				for _, prefix := range result.Matching {
					sendTranscriptLine(prefix)
				}
				for _, prefix := range result.Post {
					sendTranscriptLine(prefix)
				}

				printedResults++
				if printedResults > 20 {
					return
				}
			}
		}
	}
}

func handleTextLine(text string) string {
	matches := paralinguisticRegex.FindAllStringIndex(text, -1)

	for _, match := range matches {
		text = text[:match[0]] + "_" + text[match[0]:match[1]] + "_" + text[match[1]:]
	}

	return text
}
