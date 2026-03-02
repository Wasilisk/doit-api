package config

import (
	"log"
	"os"

	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() *Config {
	root, _ := filepath.Abs(filepath.Join(".", "..", ".."))

	err := godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		log.Println(".env file not found, using environment variables")
	}

	cfg := &Config{
		Port: getEnv("PORT", "8080"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
