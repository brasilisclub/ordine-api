package cmd

import (
	"ordine-api/pkg/routes"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	routes.ProductRoutes(r)
	routes.OrdineRoutes(r)

	r.Run()
}
