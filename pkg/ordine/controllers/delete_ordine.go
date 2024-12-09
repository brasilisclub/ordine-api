package controllers

import (
	"fmt"
	"net/http"
	"ordine-api/pkg/ordine/services"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// DeleteOrdine godoc
//
// @Summary		Delete ordine
// @Description	Delete an ordine based on the given ID
// @Tags			Ordine
// @Accept			json
// @Produce			json
// @Param			id	path		string	true	"Ordine id"
// @Success		200		{object}	utils.GenericResponse "Resource deleted successfully"
// @Failure		400		{object}	utils.GenericResponse "Error trying to delete ordine"
// @Router			/ordines/{id} [delete]
func DeleteOrdine(ctx *gin.Context) {
	err := services.RemoveOrdine(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenericResponse{
			Message: fmt.Sprintf("Error trying to delete ordine: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.GenericResponse{
		Message: "Resource deleted successfully",
	})
}
