package main

import (
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/wasilisk/doit-api/internal/config"
	"github.com/wasilisk/doit-api/internal/database"
)

func main() {
	root, _ := filepath.Abs(filepath.Join(".", "..", ".."))
	cfg := config.Load()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	if err := goose.Up(db, filepath.Join(root, "migrations")); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	app := NewApp(db, cfg)

	r := gin.Default()
	app.RegisterRoutes(r, cfg.JWT_SECRET)

	r.Run(":" + cfg.Port)
}
