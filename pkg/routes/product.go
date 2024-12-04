package routes

import (
	"ordine-api/pkg/product"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/product", product.GetProducts)
	r.POST("/product", product.PostProduct)
	r.PUT("/product/:id", updateProduct)
	r.DELETE("/product/:id", product.DeleteProduct)
}
