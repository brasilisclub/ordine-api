package controllers

import (
	"fmt"
	"net/http"
	ord "ordine-api/pkg/ordine"
	"ordine-api/pkg/ordine/services"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// PutOrdine godoc
//
// @Summary		Update an ordine
// @Description	Update an existing ordine by its ID
// @Tags			Ordine
// @Accept			json
// @Produce			json
// @Param			id	path		string	true	"Ordine ID"
// @Param			body	body		ord.OrdineRequestBody	true	"Ordine data"
// @Success		200		{object}	ordine.Ordine "Updated ordine"  // Retorna o ordine atualizado
// @Failure		400		{object}	utils.GenericResponse "Invalid input or error updating ordine"
// @Failure		500		{object}	utils.GenericResponse "Internal server error"
// @Router			/ordine/{id} [put]
func PutOrdine(ctx *gin.Context) {
	var ordine ord.OrdineRequestBody
	err := ctx.ShouldBindBodyWithJSON(&ordine)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenericResponse{
			Message: fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}
	updatedOrdine, err := services.UpdateOrdine(ctx.Param("id"), &ordine)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.GenericResponse{
			Message: fmt.Sprintf("Error trying update ordine: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedOrdine)

}
