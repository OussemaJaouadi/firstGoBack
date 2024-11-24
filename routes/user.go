package routes

import (
	"go-feToDo/controllers"
	"go-feToDo/middleware"

	"github.com/gin-gonic/gin"
)

// UserRoutes sets up user-related routes
func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	userGroup.Use(middleware.IsAuthenticated())
	{
		userGroup.GET("/", controllers.GetUserById)
		userGroup.PUT("/", controllers.UpdateUser)
		userGroup.DELETE("/", controllers.DeleteUser)
	}
}
