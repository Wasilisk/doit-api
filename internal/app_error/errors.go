package apperror

type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code ErrorCode, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

var (
	ErrEmailAlreadyExists    = New(CodeEmailAlreadyExists, "Email already exists")
	ErrInvalidCredentials    = New(CodeInvalidCredentials, "Invalid credentials")
	ErrPasswordHashingFailed = New(CodePasswordHashingFailed, "Password hashing failed")
	ErrProfileCreationFailed = New(CodeProfileCreationFailed, "Profile creation failed")

	ErrNotFound     = New(CodeNotFound, "Not found")
	ErrInternal     = New(CodeInternal, "Internal server error")
	ErrUnauthorized = New(CodeUnauthorized, "Unauthorized")
	ErrValidation   = New(CodeValidation, "Validation error")
)
