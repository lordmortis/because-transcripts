package httpServer

import (
	"BecauseLanguageBot/datasource"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
	"os"
)

func handleEpisode(ctx *gin.Context) {
	dataSource := datasource.GetSourceFromContext(ctx)
	episodeID := datasource.UUIDFromString(ctx.Param("id"))
	if episodeID == uuid.Nil {
		ctx.HTML(http.StatusBadRequest, "episodeError.html", gin.H{"error": "unable to parse id"})
		return
	}

	episodeModel, err := dataSource.EpisodeWithId(ctx, episodeID)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Couldn't fetch episode - server Error: %s\n", err))
		ctx.HTML(http.StatusInternalServerError, "episodeError.html", gin.H{"error": "server error"})
		return
	}

	if episodeModel == nil {
		ctx.HTML(http.StatusNotFound, "episodeError.html", gin.H{"error": "episode not found"})
		return
	}

	turns, count, err := episodeModel.Turns(ctx, 0, 0, true, true)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Couldn't fetch turns - server Error: %s\n", err))
		ctx.HTML(http.StatusInternalServerError, "episodeError.html", gin.H{"error": "server error"})
		return
	}

	ctx.HTML(http.StatusOK, "episode.html", gin.H{"episode": episodeModel, "turns": turns, "turnCount": count})
}
