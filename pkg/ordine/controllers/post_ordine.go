package controllers

import (
	"fmt"
	"net/http"
	ord "ordine-api/pkg/ordine"
	"ordine-api/pkg/ordine/services"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// PostOrdine godoc
//
// @Summary		Create a new ordine
// @Description	Create a new ordine with the provided data
// @Tags			Ordine
// @Accept			json
// @Produce			json
// @Param			body	body		ord.OrdineRequestBody	true	"Ordine data"
// @Success		200		{object}	ordine.Ordine "Created ordine"
// @Failure		400		{object}	utils.GenericResponse "Invalid input or error creating ordine"
// @Router			/ordines [post]
func PostOrdine(ctx *gin.Context) {
	var ordine ord.OrdineRequestBody

	err := ctx.ShouldBindBodyWithJSON(&ordine)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenericResponse{
			Message: fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}
	result, err := services.CreateOrdine(&ordine)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenericResponse{
			Message: fmt.Sprintf("Error trying to create ordine: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
