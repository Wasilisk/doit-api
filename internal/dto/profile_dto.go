package dto

type UpdateProfileRequest struct {
	FullName *string `form:"full_name"`
}

type ProfileResponse struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Email     string  `json:"email"`
	FullName  string  `json:"full_name"`
	AvatarURL *string `json:"avatar_url"`
}
