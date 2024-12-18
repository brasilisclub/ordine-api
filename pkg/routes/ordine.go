package routes

import (
	middleware "ordine-api/pkg/auth/middlewares"
	"ordine-api/pkg/ordine/controllers"

	"github.com/gin-gonic/gin"
)

func OrdineRoutes(r *gin.Engine) {

	protected := r.Group("/protected")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/ordines", controllers.GetOrdines)
		protected.GET("/ordines/:id", controllers.GetOrdine)
		protected.POST("/ordines", controllers.PostOrdine)
		protected.POST("/ordines/:id/products", controllers.PostProductsToOrdine)
		protected.PUT("/ordines/:id", controllers.PutOrdine)
		protected.DELETE("/ordines/:id", controllers.DeleteOrdine)
	}
}
