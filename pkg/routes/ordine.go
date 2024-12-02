package routes

import (
	"github.com/gin-gonic/gin"
)

func OrdineRoutes(r *gin.Engine) {
	r.GET("/ordine", getAllOrdines)
	r.POST("/ordine", createOrdine)
	r.PUT("/ordine/:id", updateOrdine)
	r.DELETE("/ordine/:id", deleteOrdine)
}
