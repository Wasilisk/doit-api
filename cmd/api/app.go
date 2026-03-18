package main

import (
	"database/sql"

	"github.com/wasilisk/doit-api/internal/config"
	"github.com/wasilisk/doit-api/internal/handler"
	"github.com/wasilisk/doit-api/internal/repository"
	"github.com/wasilisk/doit-api/internal/service"
	"github.com/wasilisk/doit-api/internal/storage"
)

type App struct {
	authHandler    *handler.AuthHandler
	profileHandler *handler.ProfileHandler
	tagHandler     *handler.TagHandler
	taskHandler    *handler.TaskHandler
}

func NewApp(db *sql.DB, cfg *config.Config) *App {
	// storage
	avatarStorage := storage.NewAvatarStorage("./static/avatars", cfg.ServerBaseURL)

	// repositories
	userRepo := repository.NewUserRepository(db)
	profileRepo := repository.NewProfileRepository(db)
	tagRepo := repository.NewTagRepository(db)
	taskRepo := repository.NewTaskRepository(db)

	// services
	authService := service.NewAuthService(userRepo, profileRepo, cfg.JWT_SECRET)
	profileService := service.NewProfileService(profileRepo, avatarStorage)
	tagService := service.NewTagService(tagRepo)
	taskService := service.NewTaskService(taskRepo)

	// handlers
	return &App{
		authHandler:    handler.NewAuthHandler(authService),
		profileHandler: handler.NewProfileHandler(profileService),
		tagHandler:     handler.NewTagHandler(tagService),
		taskHandler:    handler.NewTaskHandler(taskService),
	}
}
