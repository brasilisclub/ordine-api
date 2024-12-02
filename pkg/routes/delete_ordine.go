package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func deleteOrdine(ctx *gin.Context) {
	ctx.String(http.StatusOK, ctx.Param("id"))
}
