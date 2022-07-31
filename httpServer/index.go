package httpServer

import (
	"BecauseLanguageBot/datasource"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleIndex(ctx *gin.Context) {
	/*	episodes, _, err := datasource.EpisodesAll(100, 0)
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Unable to read episodes: %s", err))
			ctx.AbortWithError(http.StatusInternalServerError, errors.New("unable to read episodes"))
			return
		}*/

	episodes := make([]datasource.Episode, 0)

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
