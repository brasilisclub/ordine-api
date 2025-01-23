package controllers

import (
	"fmt"
	"net/http"
	prod "ordine-api/pkg/product"
	"ordine-api/pkg/product/services"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// PutProduct godoc
//
// @Summary		Update an existing product
// @Description	Update a product based on the given ID and provided details
// @Tags			Product
// @Accept			json
// @Produce			json
// @Param			id		path		string					true	"Product ID"
// @Param			product	body		product.ProductRequestBody	true	"Product data"
// @Success		200		{object}	product.Product
// @Failure		400		{object}	utils.GenericResponse "Invalid body or error updating product"
// @Failure		500		{object}	utils.GenericResponse "Internal server error"
// @Router			/products/{id} [put]
func PutProduct(ctx *gin.Context) {
	var product prod.ProductRequestBody
	if err := ctx.ShouldBindBodyWithJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenericResponse{
			Message: fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}

	updatedProduct, err := services.UpdateProduct(ctx.Param("id"), &product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.GenericResponse{
			Message: fmt.Sprintf("Error trying update product: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}
