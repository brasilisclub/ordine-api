package http

import (
	"ordine-api/api"
	"ordine-api/config"
	"ordine-api/pkg/http/routes"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetServer() *gin.Engine {
	r := gin.Default()
	r.Use(config.CorsMiddleware())

	routes.Load(r)
	//docs.SwaggerInfo.BasePath = "/api/v1"
	api.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
