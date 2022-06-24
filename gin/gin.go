package ginredoc

import (
	"github.com/LcKoh0830/go-redoc"
	"github.com/gin-gonic/gin"
)

func New(doc redoc.Redoc) gin.HandlerFunc {
	handle := doc.Handler()
	return func(ctx *gin.Context) {
		handle(ctx.Writer, ctx.Request)
	}
}
