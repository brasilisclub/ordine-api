package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"ordine-api/pkg/ordine"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/product", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"product1": "example",
		})
	})
	r.POST("/product", func(ctx *gin.Context) {
		defer ctx.Request.Body.Close()
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.String(http.StatusBadRequest, "Error reading body request")
		}
		productBody := ordine.Product{}
		err = json.Unmarshal(body, &productBody)
		if err != nil {
			ctx.String(http.StatusBadRequest, "Error trying unmarshal request body")
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"name": productBody.Name,
		})
	})
	r.PUT("/product/:id", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, ctx.Param("id"))
	})
	r.DELETE("/product/:id", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, ctx.Param("id"))
	})
}
