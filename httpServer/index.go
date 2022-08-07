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
	dataSource := datasource.GetSourceFromContext(ctx)
	episodes, _, err := dataSource.EpisodesAll(ctx, -1, -1, false)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Unable to read episodes: %s", err))
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("unable to read episodes"))
		return
	}

	data := gin.H{"episodes": episodes}

	ctx.HTML(http.StatusOK, "index.html", data)
}
