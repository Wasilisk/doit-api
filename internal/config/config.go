package config

import (
	"log"
	"os"

	"path/filepath"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type Config struct {
	Port       string
	DB         DBConfig
	JWT_SECRET string
}

func Load() *Config {
	root, _ := filepath.Abs(filepath.Join(".", "..", ".."))

	err := godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		log.Println(".env file not found, using environment variables")
	}

	cfg := &Config{
		Port: getEnv("PORT", "8080"),
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "postgres"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
