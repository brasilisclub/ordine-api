package routes

import (
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/product", getAllProducts)
	r.POST("/product", createProduct)
	r.PUT("/product/:id", updateProduct)
	r.DELETE("/product/:id", deleteProduct)
}
