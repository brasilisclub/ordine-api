package ordine

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostProductsToOrdine(ctx *gin.Context) {
	var body []OrderProductBody

	err := ctx.Bind(&body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}

	updatedOrdine, err := addProductsToOrdine(ctx.Param("id"), body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error trying update ordine: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedOrdine)

}
