package routes

import (
	"ordine-api/pkg/product"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/product", getAllProducts)
	r.POST("/product", product.PostProduct)
	r.PUT("/product/:id", updateProduct)
	r.DELETE("/product/:id", deleteProduct)
}
