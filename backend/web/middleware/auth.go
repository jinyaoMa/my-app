package middleware

import (
	"log"
	"my-app/backend/app"
	"my-app/backend/model"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		app.App()
		log.Println(model.MyOption{})
		ctx.Next()
	}
}
