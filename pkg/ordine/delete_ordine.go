package ordine

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOrdine(ctx *gin.Context) {
	err := removeOrdine(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error trying to delete ordine: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Resource deleted successfully",
	})
}
