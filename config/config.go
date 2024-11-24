package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

// Config structure holds all the configuration for the application
type Config struct {
	Env           string
	AppPort       string
	DbHost        string
	DbPort        int
	DbUser        string
	DbPassword    string
	DbName        string
	JwtSecret     string
	RefreshSecret string
	TokenExpiry   int
	RefreshExpiry int
}

var (
	cfg  *Config
	once sync.Once
)

// LoadConfig loads configuration variables from .env and environment variables
func LoadConfig() *Config {
	once.Do(func() {
		// Load .env file if it exists
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, relying on environment variables")
		}

		// Parse and validate configuration variables
		cfg = &Config{
			Env:           getEnv("ENV", "production"),
			AppPort:       getEnv("APP_PORT", "8080"),
			DbHost:        getEnv("DB_HOST", "localhost"),
			DbPort:        getEnvAsInt("DB_PORT", 5432),
			DbUser:        getEnv("DB_USER", "postgres"),
			DbPassword:    getEnv("DB_PASSWORD", "password"),
			DbName:        getEnv("DB_NAME", "echo_app"),
			JwtSecret:     getEnv("JWT_SECRET", "supersecretkey"),
			RefreshSecret: getEnv("REFRESH_SECRET", "superrefreshkey"),
			TokenExpiry:   getEnvAsInt("TOKEN_EXPIRY", 3600),    // Default: 1 hour
			RefreshExpiry: getEnvAsInt("REFRESH_EXPIRY", 86400), // Default: 1 day
		}
	})

	return cfg
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt gets an environment variable as an integer or returns a default value
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		c.DbHost,
		c.DbUser,
		c.DbPassword,
		c.DbName,
		c.DbPort,
	)
}
func (c *Config) GetSecret() string {
	return cfg.JwtSecret
}
