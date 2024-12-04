package product

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostProduct(ctx *gin.Context) {
	var product Product
	err := ctx.Bind(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}

	err = createProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Internal Server error: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}
