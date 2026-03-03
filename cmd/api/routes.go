package main

import (
	"github.com/gin-gonic/gin"
)

func (a *App) RegisterRoutes(r *gin.Engine, jwtSecret string) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", a.authHandler.Register)
		auth.POST("/login", a.authHandler.Login)
	}
}
