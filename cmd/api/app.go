package main

import (
	"database/sql"

	"github.com/wasilisk/doit-api/internal/config"
	"github.com/wasilisk/doit-api/internal/handler"
	"github.com/wasilisk/doit-api/internal/repository"
	"github.com/wasilisk/doit-api/internal/service"
)

type App struct {
	authHandler    *handler.AuthHandler
	profileHandler *handler.ProfileHandler
}

func NewApp(db *sql.DB, cfg *config.Config) *App {
	// repositories
	userRepo := repository.NewUserRepository(db)
	profileRepo := repository.NewProfileRepository(db)

	// services
	authService := service.NewAuthService(userRepo, profileRepo, cfg.JWT_SECRET)
	profileService := service.NewProfileService(profileRepo)

	// handlers
	return &App{
		authHandler:    handler.NewAuthHandler(authService),
		profileHandler: handler.NewProfileHandler(profileService),
	}
}
