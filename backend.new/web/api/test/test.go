package test

import (
	"my-app/backend.new/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Test
// @Description  Test pass path
// @Tags         Test
// @Accept       json
// @Produce      json
// @Success      200 {string} string "Pass Path"
// @Router       /test [get]
func Test() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		app.App().Log().Web()
		ctx.String(http.StatusOK, "Pass "+ctx.Request.URL.Path)
	}
}
