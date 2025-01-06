package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

// Config represents the application configuration
type Config struct {
	UpdateInterval time.Duration
	Currency       string
	AlertThreshold float64
	APIKey         string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	return Config{
		UpdateInterval: getEnvDuration("UPDATE_INTERVAL", 10*time.Second),
		Currency:       getEnvString("CURRENCY", "usd"),
		AlertThreshold: getEnvFloat("ALERT_THRESHOLD", 0),
	}
}

func getEnvString(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return fallback
}

func getEnvFloat(key string, fallback float64) float64 {
	if value, exists := os.LookupEnv(key); exists {
		if f, err := strconv.ParseFloat(value, 64); err == nil {
			return f
		}
	}
	return fallback
}
