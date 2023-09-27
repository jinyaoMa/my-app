package api

import (
	"my-app/backend/api/user"
	_ "my-app/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title My App (backend/api)
// @version 0.0.0
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

func Setup() func(app *fiber.App) *fiber.App {
	return func(app *fiber.App) *fiber.App {
		app.Get("/swagger/*", swagger.New(swagger.Config{
			// URL: "http://example.com/doc.json",
			// DeepLinking: false,
			// // Expand ("list") or Collapse ("none") tag groups by default
			// DocExpansion: "none",
			// // Prefill OAuth ClientId on Authorize popup
			// OAuth: &swagger.OAuthConfig{
			// 	AppName:  "OAuth Provider",
			// 	ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
			// },
			// // Ability to change OAuth2 redirect uri location
			// OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
		}))

		api := app.Group("/api")
		{
			user.Apply(api)
		}

		return app
	}
}
