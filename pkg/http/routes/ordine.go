package routes

import (
	"ordine-api/pkg/ordine/controllers"

	"github.com/gin-gonic/gin"
)

func OrdineRoutes(r *gin.Engine) {
	r.GET("/ordines", controllers.GetOrdines)
	r.GET("/ordine/:id", controllers.GetOrdine)
	r.POST("/ordine", controllers.PostOrdine)
	r.PUT("/ordine/:id", controllers.PutOrdine)
	r.DELETE("/ordine/:id", controllers.DeleteOrdine)
}
