package dto

import "time"

type CreateTagRequest struct {
	Name  string `json:"name"  binding:"required"`
	Color string `json:"color" binding:"required"`
}

type UpdateTagRequest struct {
	Name  *string `json:"name"`
	Color *string `json:"color"`
}

type TagResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	TaskCount int       `json:"task_count"`
	CreatedAt time.Time `json:"created_at"`
}
