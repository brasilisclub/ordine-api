package routes

import (
	"ordine-api/pkg/auth/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	//	@BasePath	/api/v1

	// PingExample godoc
	//	@Summary	ping example
	//	@Schemes
	//	@Description	do ping
	//	@Tags			example
	//	@Accept			json
	//	@Produce		json
	//	@Success		200	{string}	Helloworld
	//	@Router			/example/helloworld [get]
	r.POST("/auth/register", controllers.PostRegister)
	r.POST("/auth/login", controllers.PostLogin)
}
