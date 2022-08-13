package httpServer

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"

	"BecauseLanguageBot/datasource"
)

func handleSpeaker(ctx *gin.Context) {
	dataSource := datasource.GetSourceFromContext(ctx)
	speakerID := datasource.UUIDFromString(ctx.Param("id"))
	if speakerID == uuid.Nil {
		ctx.HTML(http.StatusBadRequest, "speakerError.html", gin.H{"error": "unable to parse id"})
		return
	}

	speakerModel, err := dataSource.SpeakerWithId(ctx, speakerID)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Couldn't fetch speaker - server Error: %s\n", err))
		ctx.HTML(http.StatusInternalServerError, "speakerError.html", gin.H{"error": "server error"})
		return
	}

	if speakerModel == nil {
		ctx.HTML(http.StatusNotFound, "speakerError.html", gin.H{"error": "speaker not found"})
		return
	}

	episodes, count, err := speakerModel.Episodes(ctx, 0, 0, false)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Couldn't fetch episodes - server Error: %s\n", err))
		ctx.HTML(http.StatusInternalServerError, "speakerError.html", gin.H{"error": "server error"})
		return
	}

	ctx.HTML(http.StatusOK, "speaker.html", gin.H{"speaker": speakerModel, "episodes": episodes, "episodeCount": count})
}
