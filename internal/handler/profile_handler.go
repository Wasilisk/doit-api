package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wasilisk/doit-api/internal/dto"
	"github.com/wasilisk/doit-api/internal/service"
	handlerutils "github.com/wasilisk/doit-api/internal/utils/handler"
)

type ProfileHandler struct {
	profileService *service.ProfileService
}

func NewProfileHandler(profileService *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService: profileService}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	profile, err := h.profileService.GetProfile(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "profile not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	req, ok := handlerutils.Bind[dto.UpdateProfileRequest](c)
	if !ok {
		return
	}

	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse form"})
		return
	}

	var avatarURL *string
	file, header, err := c.Request.FormFile("avatar")
	if err == nil {
		defer file.Close()

		url, err := h.profileService.UploadAvatar(c.Request.Context(), userID, file, header)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		avatarURL = &url
	}

	updateInput := service.UpdateProfileInput{
		UserID:    userID,
		FullName:  req.FullName,
		AvatarURL: avatarURL,
	}

	profile, err := h.profileService.UpdateProfile(c.Request.Context(), updateInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
