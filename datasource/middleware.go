package datasource

import (
	"github.com/gin-gonic/gin"
)

const contextName = "DataSource"

func (source *DataSource) Middleware() (gin.HandlerFunc, error) {
	handler := func(ctx *gin.Context) {
		ctx.Set(contextName, source)
	}

	return handler, nil
}

func GetSourceFromContext(ctx *gin.Context) (source *DataSource) {
	value := ctx.MustGet(contextName)

	source, ok := value.(*DataSource)
	if !ok {
		panic("could not get datasource from gin context")
	}

	return source
}
