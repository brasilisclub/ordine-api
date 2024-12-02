package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func updateOrdine(ctx *gin.Context) {
	ctx.String(http.StatusOK, ctx.Param("id"))
}
