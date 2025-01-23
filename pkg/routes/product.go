package routes

import (
	"ordine-api/pkg/product/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetProducts)
	r.GET("/product/:id", controllers.GetProduct)
	r.POST("/product", controllers.PostProduct)
	r.PUT("/product/:id", controllers.PutProduct)
	r.DELETE("/product/:id", controllers.DeleteProduct)
}
