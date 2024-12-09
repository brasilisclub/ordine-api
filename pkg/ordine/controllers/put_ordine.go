package controllers

import (
	"fmt"
	"net/http"
	ord "ordine-api/pkg/ordine"
	"ordine-api/pkg/ordine/services"

	"github.com/gin-gonic/gin"
)

func PutOrdine(ctx *gin.Context) {
	var ordine ord.Ordine
	err := ctx.Bind(&ordine)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}
	updatedOrdine, err := services.UpdateOrdine(ctx.Param("id"), &ordine)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error trying update ordine: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedOrdine)

}
