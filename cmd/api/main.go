package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wasilisk/doit-api/internal/config"
)

func main() {
	cfg := config.Load()

	r := gin.Default()

	r.Run(":" + cfg.Port)
}
