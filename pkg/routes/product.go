package routes

import (
	"net/http"
	"ordine-api/pkg/ordine"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/product", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"product1": "example",
		})
	})
	r.POST("/product", func(ctx *gin.Context) {
		product := ordine.Product{}

		defer ctx.Request.Body.Close()
		err := utils.UnmarshalBody(ctx.Request.Body, &product)

		if err != nil {
			ctx.String(http.StatusBadRequest, "Error trying unmarshal request body")
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"name": product.Name,
		})
	})
	r.PUT("/product/:id", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, ctx.Param("id"))
	})
	r.DELETE("/product/:id", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, ctx.Param("id"))
	})
}
