package routes

import (
	"github.com/gin-gonic/gin"
	"gotask-pro/controllers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
	}
}
