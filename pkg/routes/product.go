package routes

import (
	middleware "ordine-api/pkg/auth/middlewares"
	"ordine-api/pkg/product/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	protected := r.Group("/protected")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/products", controllers.GetProducts)
		protected.GET("/products/:id", controllers.GetProduct)
		protected.POST("/products", controllers.PostProduct)
		protected.PUT("/products/:id", controllers.PutProduct)
		protected.DELETE("/products/:id", controllers.DeleteProduct)
	}
}
