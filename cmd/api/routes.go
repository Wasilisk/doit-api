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
	}
	r.Static("/static", "./static")
}
