package routes

import (
	"ordine-api/pkg/ordine"

	"github.com/gin-gonic/gin"
)

func OrdineRoutes(r *gin.Engine) {
	r.GET("/ordine", ordine.GetOrdines)
	r.GET("/ordine/:id", ordine.GetOrdine)
	r.POST("/ordine", ordine.PostOrdine)
	r.POST("/ordine/product/:id", ordine.PostProductsToOrdine)
	r.PUT("/ordine/:id", ordine.PutOrdine)
	r.DELETE("/ordine/:id", ordine.DeleteOrdine)
}
