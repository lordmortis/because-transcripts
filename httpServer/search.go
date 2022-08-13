package httpServer

import (
	"BecauseLanguageBot/datasource"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type SearchParams struct {
	SearchString string `form:"searchString"`
}

type EpisodeMatch struct {
	Episode    *datasource.Episode
	Utterances []*datasource.Utterance
}

func handleSearch(ctx *gin.Context) {
	dataSource := datasource.GetSourceFromContext(ctx)

	var searchParams SearchParams
	err := ctx.BindQuery(&searchParams)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	utterances, count, err := dataSource.UtterancesWithText(
		ctx, searchParams.SearchString, -1, -1, true, true, true)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Couldn't search utterances - server Error: %s\n", err))
		ctx.HTML(http.StatusInternalServerError, "searchError.html", gin.H{"error": "server error"})
		return
	}

	episodeMatches := make(map[string]EpisodeMatch)

	for _, utterance := range utterances {
		episodeMatch, ok := episodeMatches[utterance.Turn.Episode.ID]
		if !ok {
			episodeMatch = EpisodeMatch{Episode: utterance.Turn.Episode}
		}
		episodeMatch.Utterances = append(episodeMatch.Utterances, utterance)
		episodeMatches[utterance.Turn.Episode.ID] = episodeMatch
	}

	data := gin.H{"episodeMatches": episodeMatches, "search": searchParams.SearchString, "utteranceCount": count, "episodeCount": len(episodeMatches)}

	ctx.HTML(http.StatusOK, "search.html", data)
}
