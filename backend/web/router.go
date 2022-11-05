package web

import (
	"embed"
	"io/fs"
	"my-app/backend/app"
	"my-app/backend/app/types"
	"my-app/backend/utils"
	"my-app/backend/web/api"
	_ "my-app/backend/web/swagger"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:embed icons
var icons embed.FS

//go:embed docs
var docs embed.FS

// router get handler for https server of web service
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
				if any == "/" {
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
		dirDocs := app.App().Cfg().Get(types.ConfigNameDirDocs)
		if utils.Utils().HasDir(dirDocs) {
			r.Static("/docs", dirDocs)
			app.App().Log().Web().Printf("WEB SERVICE SERVES DOCS FROM dirDocs: %s\n", dirDocs)
		} else {
			sub, _ := fs.Sub(docs, "docs")
			r.StaticFS("/docs", http.FS(sub))
			// extract docs into dirDocs
			assetHelper := utils.NewEmbedFS(docs, "docs")
			if err := assetHelper.Extract(dirDocs); err != nil {
				app.App().Log().Web().Println("WEB SERVICE SERVES DOCS FROM embed: backend/web/docs")
			}
		}
	}

	api.UseAPI(r)

	return r
}
