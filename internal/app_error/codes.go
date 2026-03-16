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
	CodeBadRequest ErrorCode = "BAD_REQUEST"

	// resource
	CodeTagAlreadyExists ErrorCode = "TAG_ALREADY_EXISTS"
	CodeTagNotFound      ErrorCode = "TAG_NOT_FOUND"
	CodeTaskNotFound     ErrorCode = "TASK_NOT_FOUND"
	CodeProfileNotFound  ErrorCode = "PROFILE_NOT_FOUND"

	// id
	CodeInvalidID ErrorCode = "INVALID_ID"

	// file
	CodeFileTypeNotAllowed ErrorCode = "FILE_TYPE_NOT_ALLOWED"
	CodeAvatarUploadFailed ErrorCode = "AVATAR_UPLOAD_FAILED"

	// form
	CodeFormParseFailed ErrorCode = "FORM_PARSE_FAILED"
)
