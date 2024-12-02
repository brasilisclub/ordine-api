package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllOrdines(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,
		gin.H{
			"name": "test",
		})
}
