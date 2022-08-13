package main

import (
	"BecauseLanguageBot/datasource"
	"BecauseLanguageBot/discord"
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strings"
)

func registerTranscriptSearch(dataSource *datasource.DataSource, baseURL string) {
	discord.RegisterCommand("searchTranscripts", "Search transcripts for a given word or phrase", handleTranscriptSearch(dataSource, baseURL))
}

var (
	paralinguisticRegex *regexp.Regexp
)

func init() {
	paralinguisticRegex = regexp.MustCompile(`\[.*\]`)
}

type episodeMatch struct {
	episode    *datasource.Episode
	utterances []*datasource.Utterance
}

func handleTranscriptSearch(dataSource *datasource.DataSource, baseURL string) discord.Command {
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

		ctx := context.Background()

		utterances, count, err := dataSource.UtterancesWithText(
			ctx,
			fmt.Sprintf(" %s ", parameters),
			5,
			0,
			true,
			true,
			true,
		)

		if err != nil {
			session.ChannelMessageSend(channel.ID, "error when searching for results")
			discord.LogChatError(message.Author, "error searching for results", err)
			return
		}

		if count == 0 {
			session.ChannelMessageSend(channel.ID, "No Results")
			return
		}

		if count < 5 {
			session.ChannelMessageSend(channel.ID, fmt.Sprintf("%d results:", count))
		} else {
			webSearch := fmt.Sprintf("%s/search?searchString=%s", baseURL, parameters)
			session.ChannelMessageSend(channel.ID, fmt.Sprintf("%d results, showing first 5:\n to see all results visit %s", count, webSearch))
		}

		episodeMatches := make(map[string]episodeMatch)
		for _, utterance := range utterances {
			aMatch, ok := episodeMatches[utterance.Turn.Episode.ID]
			if !ok {
				aMatch = episodeMatch{episode: utterance.Turn.Episode}
			}
			aMatch.utterances = append(aMatch.utterances, utterance)
			episodeMatches[utterance.Turn.Episode.ID] = aMatch
		}

		for _, match := range episodeMatches {
			episodeLink := fmt.Sprintf("%s/episode/%s", baseURL, match.episode.ID)
			msg := ""
			if len(match.episode.Name) > 0 {
				msg = fmt.Sprintf("Episode %d _%s_ (%s):", match.episode.Number, match.episode.Name, episodeLink)
			} else {
				msg = fmt.Sprintf("Episode %d (%s):", match.episode.Number, episodeLink)
			}
			session.ChannelMessageSend(channel.ID, msg)
			for _, utterance := range match.utterances {
				utteranceLink := fmt.Sprintf("%s/episode/%s#utterance_%s", baseURL, match.episode.ID, utterance.ID)
				msg := ""
				if len(utterance.Speakers) > 0 {
					speakerStrings := make([]string, len(utterance.Speakers))
					for index, speaker := range utterance.Speakers {
						speakerStrings[index] = fmt.Sprintf("**%s**", speaker.TranscriptName)
					}
					msg += strings.Join(speakerStrings, ",") + ": "
				}
				msg += utterance.Utterance
				msg += fmt.Sprintf("\n(context: %s)", utteranceLink)
				session.ChannelMessageSend(channel.ID, msg)
			}
		}
	}
}
