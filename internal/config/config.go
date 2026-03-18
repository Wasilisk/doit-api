package config

import (
	"log"
	"os"

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
	Port          string
	DB            DBConfig
	JWT_SECRET    string
	ClientOrigin  string
	ServerBaseURL string
}

func Load() *Config {
	err := godotenv.Load(".env")
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
		ClientOrigin:  getEnv("CLIENT_ORIGIN", "http://localhost:3000"),
		ServerBaseURL: getEnv("SERVER_BASE_URL", "http://localhost:8080"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
