package controllers

import (
	"fmt"
	"net/http"
	"ordine-api/pkg/ordine/services"

	"github.com/gin-gonic/gin"
)

func GetOrdines(ctx *gin.Context) {
	ordines, err := services.GetAllOrdines()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("An unexpected error occurred while processing the request: %s", err.Error()),
		})
	}

	ctx.JSON(http.StatusOK, ordines)
}
