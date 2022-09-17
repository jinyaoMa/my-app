package test

import (
	"my-app/backend/app"
	"my-app/backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Pass
// @Description  Test pass path
// @Tags         Test
// @Accept       json
// @Produce      json
// @Success      200 {string} string "Pass Path"
// @Router       /test [get]
func Test() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		app.App().WebLog().Println(model.MyOption{})
		ctx.String(http.StatusOK, "Pass "+ctx.Request.URL.Path)
	}
}
