package database

import (
	"log"
	"sync"

	"go-feToDo/config"
	"go-feToDo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Connect initializes the database connection
func Connect() *gorm.DB {
	once.Do(func() {
		cfg := config.LoadConfig()

		dsn := cfg.GetDSN()

		// Connect to the database
		connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		db = connection
		log.Println("Database connection established successfully!")
	})

	return db
}

// GetDB provides the singleton instance of the database connection
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database connection is not initialized. Call Connect() first.")
	}
	return db
}
func AutoMigrate(db *gorm.DB) {
	cfg := config.LoadConfig()

	if cfg.Env == "dev" {
		log.Println("Running auto migrations for development...")

		err := db.AutoMigrate(
			&models.User{},
			&models.Todo{},
		)
		if err != nil {
			log.Fatalf("Failed to auto migrate: %v", err)
		}

		log.Println("Auto migrations completed successfully!")
	} else {
		log.Println("Skipping auto migrations in production mode.")
	}
}
