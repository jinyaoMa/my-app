package static

import (
	"embed"
	"io/fs"
	_ "my-app/backend/web/static/swagger"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:embed certs docs favicon.ico
var static embed.FS

// Get self-signed certificate and key for TLS
func Certs() (crt []byte, key []byte) {
	crt, _ = static.ReadFile("certs/localhost.crt")
	key, _ = static.ReadFile("certs/localhost.key")
	return
}

// Setup static resources and websites
func Setup(r *gin.Engine) *gin.Engine {
	// setup favicon
	r.StaticFileFS("/favicon.ico", "favicon.ico", http.FS(static))

	// setup vitepress docs
	vp, _ := fs.Sub(static, "docs")
	r.StaticFS("/docs", http.FS(vp))

	// setup swagger ui
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
