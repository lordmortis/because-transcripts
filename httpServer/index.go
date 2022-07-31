package httpServer

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/errgo.v2/errors"

	"BecauseLanguageBot/datasource"
)

func handleIndex(ctx *gin.Context) {
	episodes, _, err := datasource.EpisodesAll(ctx, 100, 0)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Unable to read episodes: %s", err))
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("unable to read episodes"))
		return
	}

	templateEpisodes := make([]gin.H, len(episodes))
	for index, episode := range episodes {
		templateEpisodes[index] = gin.H{
			"id":    episode.ID,
			"name":  episode.Name,
			"aired": episode.Date.String(),
		}
	}

	data := gin.H{"Episodes": templateEpisodes}

	ctx.HTML(http.StatusOK, "index.html", data)
}
