package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wasilisk/doit-api/internal/dto"
	handlerutils "github.com/wasilisk/doit-api/internal/handler_utils"
	"github.com/wasilisk/doit-api/internal/service"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	req, ok := handlerutils.BindJSON[dto.CreateTaskRequest](c)
	if !ok {
		return
	}

	task, err := h.taskService.CreateTask(c.Request.Context(), userID, *req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) PatchTask(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	taskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}

	req, ok := handlerutils.BindJSON[dto.UpdateTaskRequest](c)
	if !ok {
		return
	}

	task, err := h.taskService.UpdateTask(c.Request.Context(), userID, taskID, *req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	var filter dto.TaskFilterRequest
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks, err := h.taskService.GetTasks(c.Request.Context(), userID, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	taskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}

	task, err := h.taskService.GetTaskByID(c.Request.Context(), userID, taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	taskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}

	if err := h.taskService.DeleteTask(c.Request.Context(), userID, taskID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *TaskHandler) RestoreTask(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)

	taskID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
		return
	}

	if err := h.taskService.RestoreTask(c.Request.Context(), userID, taskID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task restored"})
}
