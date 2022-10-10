package static

import (
	"embed"
	"my-app/backend/pkg/utils"
	_ "my-app/backend/web/static/swagger"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:embed certs favicon.ico
var static embed.FS

// Get self-signed certificate and key for TLS
func Certs() (crt []byte, key []byte) {
	crt, _ = static.ReadFile("certs/localhost.crt")
	key, _ = static.ReadFile("certs/localhost.key")
	return
}

// setup favicon
func SetupFavicon(r *gin.Engine) *gin.Engine {
	r.StaticFileFS("/favicon.ico", "favicon.ico", http.FS(static))
	return r
}

// setup swagger ui
func SetupSwaggerUI(r *gin.Engine) *gin.Engine {
	r.GET("/swagger", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	r.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.PersistAuthorization(true),
		),
	)
	return r
}

// setup vitepress docs
func SetupVitePress(r *gin.Engine) *gin.Engine {
	r.Static("/docs", utils.GetExecutablePath("Docs")) // build/bin/Docs
	return r
}
