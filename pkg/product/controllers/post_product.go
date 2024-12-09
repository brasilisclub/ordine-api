package controllers

import (
	"fmt"
	"net/http"
	prod "ordine-api/pkg/product"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// PostProduct godoc
//
// @Summary		Create a new product
// @Description	Create a new product based on the provided data
// @Tags			Product
// @Accept			json
// @Produce			json
// @Param			product	body		product.ProductRequestBody	true		"Product data"
// @Success		201		{object}	product.Product
// @Failure		400		{object}	utils.GenericResponse "Invalid body or bad request"
// @Failure		500		{object}	utils.GenericResponse "Internal server error"
// @Router			/products [post]
func PostProduct(ctx *gin.Context) {
	var product prod.ProductRequestBody
	err := ctx.ShouldBindBodyWithJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenericResponse{
			Message: fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}

	result, err := prod.CreateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.GenericResponse{
			Message: fmt.Sprintf("Internal Server error: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
