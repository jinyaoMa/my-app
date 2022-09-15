package web

import (
	_ "my-app/backend/web/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My App (backend/web/router.go)
// @version 1.0.0
// @description "My App is a continuously updated personal service collection."

// @contact.name Github Issues
// @contact.url https://github.com/jinyaoMa/my-app/issues

// @license.name MIT
// @license.url https://github.com/jinyaoMa/my-app/blob/main/LICENSE

// @BasePath /api

// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization

func router() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Author": "pong",
		})
	})

	r.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerfiles.Handler,
			ginSwagger.PersistAuthorization(true),
		),
	)

	return r
}
