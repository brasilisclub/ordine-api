package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"product1": "example",
	})
}
