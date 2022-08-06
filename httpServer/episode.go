package httpServer

import (
	"BecauseLanguageBot/datasource"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
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
		ctx.HTML(http.StatusInternalServerError, "episodeError.html", gin.H{"error": "server error"})
		return
	}

	if episodeModel == nil {
		ctx.HTML(http.StatusNotFound, "episodeError.html", gin.H{"error": "episode not found"})
		return
	}

	utterances, count, err := episodeModel.Utterances(ctx, 0, 0, true)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "episodeError.html", gin.H{"error": "server error"})
		return
	}

	ctx.HTML(http.StatusOK, "episode.html", gin.H{"episode": episodeModel, "utterances": utterances, "count": count})
}
