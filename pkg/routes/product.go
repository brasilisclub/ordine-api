package routes

import (
	"ordine-api/pkg/product"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/product", product.GetProducts)
	r.GET("/product/:id", product.GetProduct)
	r.POST("/product", product.PostProduct)
	r.PUT("/product/:id", product.PutProduct)
	r.DELETE("/product/:id", product.DeleteProduct)
}
