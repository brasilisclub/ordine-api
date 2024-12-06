package ordine

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrdine(ctx *gin.Context) {
	ordine, err := GetOrdineById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("An unexpected error occurred while processing the request: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, ordine)
}
