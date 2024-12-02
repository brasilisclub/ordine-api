package main

import (
	"ordine-api/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.ProductRoutes(r)
	routes.OrdineRoutes(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
