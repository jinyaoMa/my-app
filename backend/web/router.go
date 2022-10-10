package web

import (
	"my-app/backend/web/api/test"
	"my-app/backend/web/middleware"
	"my-app/backend/web/static"

	"github.com/gin-gonic/gin"
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
	r := gin.Default()
	{
		static.SetupFavicon(r)
		static.SetupSwaggerUI(r)
		static.SetupVitePress(r)
	}

	a := r.Group("/auth")
	{
		a.GET("/", middleware.Auth(), func(ctx *gin.Context) {})
	}

	b := r.Group("/api")
	{
		b.GET("/test", test.Test())
	}

	return r
}
