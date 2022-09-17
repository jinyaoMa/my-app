package web

import (
	"my-app/backend/app"
	"my-app/backend/web/api/test"
	"my-app/backend/web/docs"
	"my-app/backend/web/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My App (backend/web/router.go)
// @version 1.0.0
// @description "My App is a continuously updated personal service collection."

// @contact.name GitHub Discussions
// @contact.url https://github.com/jinyaoMa/my-app/discussions

// @license.name MIT
// @license.url https://github.com/jinyaoMa/my-app/blob/main/LICENSE

// @schemes https
// @BasePath /api

// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
// @description Authorization Header should contain value started with "Bearer " and followed by a JSON Web Token.

func router() *gin.Engine {
	if app.App().Env().Log2File() {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}
	gin.DefaultWriter = app.App().WebLog().Writer()

	r := gin.Default()

	r.GET("/", test.Test())

	a := r.Group("/auth")
	{
		a.GET("/", middleware.Auth(), func(ctx *gin.Context) {})
	}

	// "/api"
	b := r.Group(docs.SwaggerInfo.BasePath)
	{
		b.GET("/test", test.Test())
	}

	r.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.PersistAuthorization(true),
		),
	)

	return r
}
