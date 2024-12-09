package controllers

import (
	"fmt"
	"net/http"
	"ordine-api/pkg/ordine/services"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// PostLogin godoc
//
//	@Summary		Delete ordine
//	@Description	Delete a ordine by id
//	@Tags			Ordine
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Ordine id"
//	@Success		200		{object}	utils.GenericResponse
//	@Failure		400		{object}	utils.GenericResponse
//	@Router			/ordines/{id} [delete]
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
