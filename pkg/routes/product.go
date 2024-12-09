package routes

import (
	"ordine-api/pkg/product/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProduct)
	r.POST("/products", controllers.PostProduct)
	r.PUT("/products/:id", controllers.PutProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
}
