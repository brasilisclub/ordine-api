package routes

import (
	"ordine-api/pkg/auth/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/auth/register", controllers.PostRegister)
	r.POST("/auth/login", controllers.PostLogin)
}
