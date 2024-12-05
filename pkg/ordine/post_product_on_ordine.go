package ordine

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostProductOnOrdine(ctx *gin.Context) {
	var body addProductsToOrdineBody

	err := ctx.Bind(&body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}

	updatedOrdine, err := addProductsToOrdine(ctx.Param("id"), body.ProductsId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error trying update ordine: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedOrdine)

}
