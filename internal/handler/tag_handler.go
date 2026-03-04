package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wasilisk/doit-api/internal/dto"
	handlerutils "github.com/wasilisk/doit-api/internal/handler_utils"
	"github.com/wasilisk/doit-api/internal/service"
)

type TagHandler struct {
	tagService *service.TagService
}

func NewTagHandler(tagService *service.TagService) *TagHandler {
	return &TagHandler{tagService: tagService}
}

func (h *TagHandler) CreateTag(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	req, ok := handlerutils.BindJSON[dto.CreateTagRequest](c)
	if !ok {
		return
	}

	tag, err := h.tagService.CreateTag(c.Request.Context(), userID, *req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tag)
}

func (h *TagHandler) GetTags(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	tags, err := h.tagService.GetTags(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tags)
}

func (h *TagHandler) UpdateTag(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	tagID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tag id"})
		return
	}

	req, ok := handlerutils.BindJSON[dto.UpdateTagRequest](c)
	if !ok {
		return
	}

	tag, err := h.tagService.UpdateTag(c.Request.Context(), userID, tagID, *req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tag)
}

func (h *TagHandler) DeleteTag(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	tagID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tag id"})
		return
	}

	if err := h.tagService.DeleteTag(c.Request.Context(), userID, tagID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
