package controllers

import (
	"fmt"
	"net/http"
	"ordine-api/pkg/ordine/services"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// GetOrdine godoc
//
// @Summary		Get ordine by ID
// @Description	Retrieve an ordine by its unique ID
// @Tags			Ordine
// @Accept			json
// @Produce			json
// @Param			id	path		string	true	"Ordine ID"
// @Success		200		{object}	ordine.Ordine "Ordine data"
// @Failure		400		{object}	utils.GenericResponse "Invalid ID format"
// @Failure		500		{object}	utils.GenericResponse "Internal server error"
// @Router			/ordines/{id} [get]
func GetOrdine(ctx *gin.Context) {
	ordine, err := services.GetOrdineById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.GenericResponse{
			Message: fmt.Sprintf("An unexpected error occurred while processing the request: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, ordine)
}
