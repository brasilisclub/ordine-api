package product

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutProduct(ctx *gin.Context) {
	var product Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}

	updatedProduct, err := updateProduct(ctx.Param("id"), &product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error trying update product: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}
