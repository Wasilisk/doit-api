package main

import (
	"database/sql"

	"github.com/wasilisk/doit-api/internal/config"
	"github.com/wasilisk/doit-api/internal/handler"
	"github.com/wasilisk/doit-api/internal/repository"
	"github.com/wasilisk/doit-api/internal/service"
)

type App struct {
	authHandler *handler.AuthHandler
}

func NewApp(db *sql.DB, cfg *config.Config) *App {
	// repositories
	userRepo := repository.NewUserRepository(db)

	// services
	authService := service.NewAuthService(userRepo, cfg.JWT_SECRET)

	// handlers
	return &App{
		authHandler: handler.NewAuthHandler(authService),
	}
}
