package controllers

import (
	"fmt"
	"net/http"
	ord "ordine-api/pkg/ordine"
	order "ordine-api/pkg/ordine/order"

	"github.com/gin-gonic/gin"
)

func PostProductsToOrdine(ctx *gin.Context) {
	var body []ord.OrderProductBody

	err := ctx.Bind(&body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}

	updatedOrdine, err := order.AddProductsToOrdine(ctx.Param("id"), body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error trying update ordine: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedOrdine)

}
