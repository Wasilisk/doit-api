package dto

type FieldError struct {
	Field string `json:"field"`
	Code  string `json:"code"`
}

type ErrorResponse struct {
	Errors []FieldError `json:"errors,omitempty"`
}
