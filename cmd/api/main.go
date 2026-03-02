package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wasilisk/doit-api/internal/config"
	"github.com/wasilisk/doit-api/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	r.Run(":" + cfg.Port)
}
