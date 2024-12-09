package controllers

import (
	"fmt"
	"net/http"
	prod "ordine-api/pkg/product"

	"github.com/gin-gonic/gin"
)

func GetProducts(ctx *gin.Context) {

	products, err := prod.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("An unexpected error occurred while processing the request: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
