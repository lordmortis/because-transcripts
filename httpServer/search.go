package httpServer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchParams struct {
	SearchString string `form:"searchString"`
}

func handleSearch(ctx *gin.Context) {
	var searchParams SearchParams
	err := ctx.BindQuery(&searchParams)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	data := gin.H{"Results": nil, "SearchString": searchParams.SearchString}

	ctx.HTML(http.StatusOK, "search.html", data)
}
