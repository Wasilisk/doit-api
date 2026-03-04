package dto

type CreateTagRequest struct {
	Name  string `json:"name"  binding:"required"`
	Color string `json:"color" binding:"required"`
}

type UpdateTagRequest struct {
	Name  *string `json:"name"`
	Color *string `json:"color"`
}

type TagResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	CreatedAt string `json:"created_at"`
}
