package controllers

import (
	"fmt"
	"net/http"
	ord "ordine-api/pkg/ordine"

	"github.com/gin-gonic/gin"
)

func DeleteOrdine(ctx *gin.Context) {
	err := ord.RemoveOrdine(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error trying to delete ordine: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Resource deleted successfully",
	})
}
