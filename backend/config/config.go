package config

import (
	"fmt"
	"os"
)

// Config holds all configuration for the application.
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	ServerPort string
}

// LoadConfig reads configuration from environment variables with sensible defaults.
func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "erp_user"),
		DBPassword: getEnv("DB_PASSWORD", "erp_password"),
		DBName:     getEnv("DB_NAME", "erp_project_mgmt"),
		JWTSecret:  getEnv("JWT_SECRET", "default-secret-key-change-in-production"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

// GetDSN returns the PostgreSQL data source name.
func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName,
	)
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
