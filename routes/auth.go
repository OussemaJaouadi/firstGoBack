package routes

import (
	"go-feToDo/controllers"

	"github.com/gin-gonic/gin"
)

// AuthRoutes defines routes related to authentication
func AuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", controllers.Login)
		authGroup.POST("/refresh", controllers.Refresh)
		authGroup.POST("/register", controllers.CreateUser)
	}
}
