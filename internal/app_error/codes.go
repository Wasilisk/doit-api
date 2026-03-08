package apperror

type ErrorCode string

const (
	// auth
	CodeEmailAlreadyExists    ErrorCode = "EMAIL_ALREADY_EXISTS"
	CodeInvalidCredentials    ErrorCode = "INVALID_CREDENTIALS"
	CodePasswordHashingFailed ErrorCode = "PASSWORD_HASHING_FAILED"
	CodeProfileCreationFailed ErrorCode = "PROFILE_CREATION_FAILED"
	CodeUnauthorized          ErrorCode = "UNAUTHORIZED"
	CodeUserWithEmailNotFound ErrorCode = "USER_WITH_EMAIL_NOT_FOUND"

	// general
	CodeNotFound   ErrorCode = "NOT_FOUND"
	CodeInternal   ErrorCode = "INTERNAL_ERROR"
	CodeValidation ErrorCode = "VALIDATION_ERROR"
)
