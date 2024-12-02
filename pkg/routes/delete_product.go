package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func deleteProduct(ctx *gin.Context) {
	ctx.String(http.StatusOK, ctx.Param("id"))
}
