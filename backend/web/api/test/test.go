package test

import (
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
		ctx.String(http.StatusOK, "Pass "+ctx.Request.URL.Path)
	}
}
