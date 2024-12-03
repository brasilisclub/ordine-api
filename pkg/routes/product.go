package routes

import (
	"ordine-api/pkg/ordine"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/product", getAllProducts)
	r.POST("/product", ordine.PostProduct)
	r.PUT("/product/:id", updateProduct)
	r.DELETE("/product/:id", deleteProduct)
}
