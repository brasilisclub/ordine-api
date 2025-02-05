package routes

import (
	"ordine-api/pkg/ordine/order/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	r.POST("/ordine/:id/products", controllers.PostProductsToOrdine)
}
