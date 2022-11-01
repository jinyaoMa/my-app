package api

import (
	"my-app/backend.new/web/api/auth"
	"my-app/backend/web/api/test"

	"github.com/gin-gonic/gin"
)

// @title My App (backend/web/api/api.go)
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

func UseAPI(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/test", test.Test())
		_auth := api.Group("/auth")
		{
			_auth.POST("/login", auth.Login())
		}
	}
}
