package env

import (
	"os"

	"github.com/joho/godotenv"
)

// Load env from a specific file
func LoadEnv(envPath string) error {
	return godotenv.Load(envPath)
}

// Get env value by key
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}