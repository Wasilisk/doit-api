package apperror

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}

type ErrorBody struct {
	Code    ErrorCode    `json:"code"`
	Message string       `json:"message"`
	Fields  []FieldError `json:"fields,omitempty"`
}

type FieldError struct {
	Field   string            `json:"field"`
	Code    ErrorCode         `json:"code"`
	Message string            `json:"message"`
	Context map[string]string `json:"context,omitempty"`
}

func HandleError(c *gin.Context, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		status := codeToStatus(appErr.Code)
		c.JSON(status, ErrorResponse{Error: ErrorBody{
			Code:    appErr.Code,
			Message: appErr.Message,
		}})
		return
	}

	c.JSON(http.StatusInternalServerError, ErrorResponse{Error: ErrorBody{
		Code:    CodeInternal,
		Message: "Something went wrong",
	}})
}

func codeToStatus(code ErrorCode) int {
	switch code {
	case CodeEmailAlreadyExists:
		return http.StatusConflict
	case CodeInvalidCredentials:
		return http.StatusUnauthorized
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeUserWithEmailNotFound:
		return http.StatusNotFound
	case CodeNotFound:
		return http.StatusNotFound
	case CodeValidation:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
