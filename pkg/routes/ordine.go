package routes

import (
	"ordine-api/pkg/ordine"

	"github.com/gin-gonic/gin"
)

func OrdineRoutes(r *gin.Engine) {
	r.GET("/ordine", getAllOrdines)
	r.POST("/ordine", ordine.PostOrdine)
	r.PUT("/ordine/:id", updateOrdine)
	r.DELETE("/ordine/:id", deleteOrdine)
}
