package web

import (
	"embed"
	"io/fs"
	"my-app/backend.new/app"
	"my-app/backend.new/model"
	"my-app/backend.new/utils"
	"my-app/backend.new/web/api"
	_ "my-app/backend.new/web/swagger"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:embed icons
var icons embed.FS

//go:embed docs
var docs embed.FS

var ()

// router get handler for https server of web server
func (w *web) router() *gin.Engine {
	r := gin.Default()
	{
		// setup favicon
		r.StaticFileFS("/favicon.ico", "icons/favicon.ico", http.FS(icons))
		// setup swagger ui
		r.GET("/swagger", func(ctx *gin.Context) {
			ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		})
		r.GET(
			"/swagger/*any",
			func(ctx *gin.Context) {
				any := ctx.Param("any")
				if any == "" {
					ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
				}
				ctx.Next()
			},
			ginSwagger.WrapHandler(
				swaggerFiles.Handler,
				ginSwagger.PersistAuthorization(true),
			),
		)
		// setup docs
		app.App().UseConfig(func(cfg *app.Config) {
			dirDocs := cfg.Get(model.OptionNameDirDocs)
			if utils.Utils().HasDir(dirDocs) {
				r.Static("/docs", dirDocs)
			} else {
				sub, _ := fs.Sub(docs, "docs")
				r.StaticFS("/docs", http.FS(sub))
				// extract docs into dirDocs
				assetHelper := utils.NewEmbedFS(docs, "docs")
				if err := assetHelper.Extract(dirDocs); err != nil {
					app.App().Log().Web().Printf("failed to extract embed docs into dirDocs (%s): %+v\n", dirDocs, err)
				}
			}
		})
	}

	api.UseAPI(r)

	return r
}
