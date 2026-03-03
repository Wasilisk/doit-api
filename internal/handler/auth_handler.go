package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wasilisk/doit-api/internal/dto"
	handlerutils "github.com/wasilisk/doit-api/internal/handler_utils"
	"github.com/wasilisk/doit-api/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	req, ok := handlerutils.BindJSON[dto.RegisterRequest](c)
	if !ok {
		return
	}

	token, err := h.authService.Register(c.Request.Context(), service.RegisterInput{
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
	})
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.AuthResponse{Token: token})
}

func (h *AuthHandler) Login(c *gin.Context) {
	req, ok := handlerutils.BindJSON[dto.LoginRequest](c)
	if !ok {
		return
	}

	token, err := h.authService.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{Token: token})
}
