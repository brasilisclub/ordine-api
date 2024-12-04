package product

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(ctx *gin.Context) {

	products, err := getAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("An unexpected error occurred while processing the request: %s", err.Error()),
		})
	}

	ctx.JSON(http.StatusOK, products)
}
