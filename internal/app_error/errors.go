package apperror

type AppError struct {
	Code ErrorCode `json:"code"`
}

func (e *AppError) Error() string {
	return string(e.Code)
}

func New(code ErrorCode) *AppError {
	return &AppError{Code: code}
}
