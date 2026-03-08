package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperror "github.com/wasilisk/doit-api/internal/app_error"
	"github.com/wasilisk/doit-api/internal/dto"
	"github.com/wasilisk/doit-api/internal/service"
	handlerutils "github.com/wasilisk/doit-api/internal/utils/handler"
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
		apperror.HandleError(c, err)
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
		apperror.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{Token: token})
}
