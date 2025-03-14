
package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Environment     string
	ServerPort      string
	DBPath          string
	JWTSecret       string
	GoogleClientID  string
	GoogleSecret    string
	GoogleRedirect  string
	TokenExpiration int // in hours
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	godotenv.Load()

	cfg := &Config{
		Environment:     getEnv("ENVIRONMENT", "development"),
		ServerPort:      getEnv("SERVER_PORT", "8080"),
		DBPath:          getEnv("DB_PATH", "data/sqlite.db"),
		JWTSecret:       getEnv("JWT_SECRET", "your-secret-key"),
		GoogleClientID:  getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleSecret:    getEnv("GOOGLE_SECRET", ""),
		GoogleRedirect:  getEnv("GOOGLE_REDIRECT", "http://localhost:8080/api/auth/google/callback"),
		TokenExpiration: getEnvAsInt("TOKEN_EXPIRATION", 24), // Default: 24 hours
	}

	return cfg, nil
}

// Helper function to get an environment variable or a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Helper function to get an environment variable as int or a default value
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
