package ordine

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutOrdine(ctx *gin.Context) {
	var ordine Ordine
	err := ctx.Bind(&ordine)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}
	updatedOrdine, err := updateOrdine(ctx.Param("id"), &ordine)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error trying update ordine: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedOrdine)

}
