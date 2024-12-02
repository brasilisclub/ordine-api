package routes

import (
	"net/http"
	"ordine-api/pkg/ordine"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func createProduct(ctx *gin.Context) {
	product := ordine.Product{}

	defer ctx.Request.Body.Close()
	err := utils.UnmarshalBody(ctx.Request.Body, &product)

	if err != nil {
		ctx.String(http.StatusBadRequest, "Error trying unmarshal request body")
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"name": product.Name,
	})
}
