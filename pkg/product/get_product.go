package product

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProduct(ctx *gin.Context) {
	product, err := getProductById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("An unexpected error occurred while processing the request: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, product)

}
