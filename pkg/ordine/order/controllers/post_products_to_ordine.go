package controllers

import (
	"fmt"
	"net/http"
	"ordine-api/pkg/ordine/order"
	"ordine-api/pkg/ordine/order/services"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// @Summary		Add products to an ordine
// @Description	Add products to an existing ordine by its ID
// @Tags			Order
// @Accept			json
// @Produce			json
// @Param			id	path		string	true	"Ordine ID"
// @Param			body	body		[]order.OrderProductBody	true	"Products to add to the ordine"
// @Success		200		{object}	ordine.Ordine "Updated ordine"
// @Failure		400		{object}	utils.GenericResponse "Invalid input or error updating ordine"
// @Failure		500		{object}	utils.GenericResponse "Internal server error"
// @Router			/ordine/{id}/products [post]
func PostProductsToOrdine(ctx *gin.Context) {
	var body []order.OrderProductBody

	err := ctx.ShouldBindBodyWithJSON(&body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenericResponse{
			Message: fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}

	updatedOrdine, err := services.AddProductsToOrdine(ctx.Param("id"), body)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.GenericResponse{
			Message: fmt.Sprintf("Error trying update ordine: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedOrdine)

}
