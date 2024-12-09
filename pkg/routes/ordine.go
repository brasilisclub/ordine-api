package routes

import (
	"ordine-api/pkg/ordine/controllers"

	"github.com/gin-gonic/gin"
)

func OrdineRoutes(r *gin.Engine) {
	r.GET("/ordines", controllers.GetOrdines)
	r.GET("/ordines/:id", controllers.GetOrdine)
	r.POST("/ordines", controllers.PostOrdine)
	r.POST("/ordines/:id/products", controllers.PostProductsToOrdine)
	r.PUT("/ordines/:id", controllers.PutOrdine)
	r.DELETE("/ordines/:id", controllers.DeleteOrdine)
}
