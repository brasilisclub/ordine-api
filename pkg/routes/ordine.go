package routes

import (
	"ordine-api/pkg/ordine"

	"github.com/gin-gonic/gin"
)

func OrdineRoutes(r *gin.Engine) {
	r.GET("/ordine", ordine.GetOrdines)
	r.GET("/ordine/:id")
	r.POST("/ordine", ordine.PostOrdine)
	r.PUT("/ordine/:id", ordine.PutOrdine)
	r.DELETE("/ordine/:id", ordine.DeleteOrdine)
}
