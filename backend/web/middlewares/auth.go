package middlewares

import (
	"my-app/backend/app"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		app.App().Log().Web()
		ctx.Next()
	}
}
