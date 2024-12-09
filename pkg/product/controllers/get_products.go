package controllers

import (
	"fmt"
	"net/http"
	prod "ordine-api/pkg/product"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// GetProducts godoc
//
// @Summary		Get all products
// @Description	Retrieve a list of all products
// @Tags			Product
// @Accept			json
// @Produce			json
// @Success		200		{array}	product.Product
// @Failure		500		{object}	utils.GenericResponse "Internal server error"
// @Router			/products [get]
func GetProducts(ctx *gin.Context) {

	products, err := prod.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.GenericResponse{
			Message: fmt.Sprintf("An unexpected error occurred while processing the request: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
