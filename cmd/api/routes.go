package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wasilisk/doit-api/internal/middleware"
)

func (a *App) RegisterRoutes(r *gin.Engine, jwtSecret string) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", a.authHandler.Register)
		auth.POST("/login", a.authHandler.Login)
	}
	api := r.Group("/api").Use(middleware.Auth(jwtSecret))
	{
		api.GET("/profile", a.profileHandler.GetProfile)
		api.PATCH("/profile", a.profileHandler.UpdateProfile)

		api.GET("/tags", a.tagHandler.GetTags)
		api.POST("/tags", a.tagHandler.CreateTag)
		api.PATCH("/tags/:id", a.tagHandler.UpdateTag)
		api.DELETE("/tags/:id", a.tagHandler.DeleteTag)

		api.GET("/tasks", a.taskHandler.GetTasks)
		api.POST("/tasks", a.taskHandler.CreateTask)
		api.PATCH("/tasks/:id", a.taskHandler.PatchTask)
		api.GET("/tasks/:id", a.taskHandler.GetTaskByID)
		api.DELETE("/tasks/:id", a.taskHandler.DeleteTask)
		api.POST("/tasks/:id/restore", a.taskHandler.RestoreTask)
	}
	r.Static("/static", "./static")
}
