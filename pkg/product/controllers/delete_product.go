package controllers

import (
	"fmt"
	"net/http"
	prod "ordine-api/pkg/product"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(ctx *gin.Context) {

	err := prod.RemoveProduct(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error trying to delete product: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Resource deleted successfully",
	})
}
