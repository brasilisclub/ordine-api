package controllers

import (
	"fmt"
	"net/http"
	ord "ordine-api/pkg/ordine"

	"github.com/gin-gonic/gin"
)

func PostOrdine(ctx *gin.Context) {
	var ordine ord.Ordine

	err := ctx.Bind(&ordine)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}
	err = ord.CreateOrdine(&ordine)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error trying to create ordine: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, ordine)
}
