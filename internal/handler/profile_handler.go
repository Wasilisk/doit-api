package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	apperror "github.com/wasilisk/doit-api/internal/app_error"
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
		apperror.HandleError(c, err)
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
		apperror.HandleError(c, apperror.New(apperror.CodeFormParseFailed))
		return
	}

	var avatarURL *string
	file, header, err := c.Request.FormFile("avatar")
	if err == nil {
		defer file.Close()

		url, err := h.profileService.UploadAvatar(c.Request.Context(), userID, file, header)
		if err != nil {
			apperror.HandleError(c, err)
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
		apperror.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, profile)
}
