package controllers

import (
	"fmt"
	"net/http"
	"ordine-api/pkg/product/services"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// GetProduct godoc
//
// @Summary		Get a product by ID
// @Description	Retrieve a product based on the given ID
// @Tags			Product
// @Accept			json
// @Produce			json
// @Param			id	path		string	true	"Product ID"
// @Success		200		{object}	product.Product
// @Failure		400		{object}	utils.GenericResponse "Invalid ID or error retrieving product"
// @Failure		500		{object}	utils.GenericResponse "Internal server error"
// @Router			/product/{id} [get]
func GetProduct(ctx *gin.Context) {
	product, err := services.GetProductById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.GenericResponse{
			Message: fmt.Sprintf("An unexpected error occurred while processing the request: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, product)

}
