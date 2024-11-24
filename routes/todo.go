package routes

import (
	"go-feToDo/controllers"
	"go-feToDo/middleware"

	"github.com/gin-gonic/gin"
)

// TodoRoutes sets up todo-related routes
func TodoRoutes(router *gin.Engine) {
	todoGroup := router.Group("/todos")
	todoGroup.Use(middleware.IsAuthenticated())
	{
		todoGroup.GET("/", controllers.GetActiveToDo)
		todoGroup.GET("/all", controllers.GetAllToDo)
		todoGroup.GET("/trash", controllers.GetTrashToDo)
		todoGroup.GET("/:todoID", controllers.GetTodoById)
		todoGroup.POST("/", controllers.CreateTodo)
		todoGroup.PUT("/:todoID", controllers.UpdateTodo)
		todoGroup.DELETE("/:todoID/trash", controllers.SoftDeleteTodo)
		todoGroup.DELETE("/:todoID/permanent", controllers.DeleteTodo)
	}
}
