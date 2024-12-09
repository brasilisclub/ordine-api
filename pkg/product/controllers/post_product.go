package controllers

import (
	"fmt"
	"net/http"
	prod "ordine-api/pkg/product"

	"github.com/gin-gonic/gin"
)

func PostProduct(ctx *gin.Context) {
	var product prod.Product
	err := ctx.ShouldBindBodyWithJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}

	err = prod.CreateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Internal Server error: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}
