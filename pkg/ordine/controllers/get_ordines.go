package controllers

import (
	"fmt"
	"net/http"
	"ordine-api/pkg/ordine/services"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

// GetOrdines godoc
//
// @Summary		Get all ordines
// @Description	Retrieve a list of all ordines
// @Tags			Ordine
// @Accept			json
// @Produce			json
// @Success		200		{array}	ordine.Ordine "List of ordines"
// @Failure		500		{object}	utils.GenericResponse "Internal server error"
// @Router			/ordines [get]
func GetOrdines(ctx *gin.Context) {
	ordines, err := services.GetAllOrdines()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.GenericResponse{
			Message: fmt.Sprintf("An unexpected error occurred while processing the request: %s", err.Error()),
		})
	}

	ctx.JSON(http.StatusOK, ordines)
}
