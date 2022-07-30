package httpServer

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
