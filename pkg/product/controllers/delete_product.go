package controllers

import (
	"fmt"
	"net/http"
	"ordine-api/pkg/product/services"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// DeleteProduct godoc
//
// @Summary		Delete a product
// @Description	Delete a product by its ID
// @Tags			Product
// @Accept			json
// @Produce			json
// @Param			id	path		string	true	"Product ID"
// @Success		200		{object}	utils.GenericResponse "Product deleted successfully"
// @Failure		400		{object}	utils.GenericResponse "Error trying to delete the product"
// @Router			/products/{id} [delete]
func DeleteProduct(ctx *gin.Context) {

	err := services.RemoveProduct(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenericResponse{
			Message: fmt.Sprintf("Error trying to delete product: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.GenericResponse{
		Message: "Resource deleted successfully",
	})
}
