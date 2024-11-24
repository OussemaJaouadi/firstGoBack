package main

import (
	"go-feToDo/config"
	"go-feToDo/database"
	"go-feToDo/routes"
	"go-feToDo/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitLogger("debug", "dev")
	// Load application configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	db := database.Connect()

	// Run database migrations if in development mode
	database.AutoMigrate(db)

	// Create a Gin router instance
	router := gin.Default()

	// Initialize application routes
	initializeRoutes(router)

	// Start the server
	log.Printf("Starting server on port %s...", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// initializeRoutes sets up all application routes
func initializeRoutes(router *gin.Engine) {
	// Route groups
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.TodoRoutes(router)
}
