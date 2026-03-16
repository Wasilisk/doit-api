package apperror

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wasilisk/doit-api/internal/i18n"
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
	Field   string    `json:"field"`
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

// getLang extracts the language from the Gin context.
func getLang(c *gin.Context) i18n.Lang {
	lang, exists := c.Get("lang")
	if !exists {
		return i18n.EN
	}
	l, ok := lang.(i18n.Lang)
	if !ok {
		return i18n.EN
	}
	return l
}

// HandleError writes a localized JSON error response based on the error type.
func HandleError(c *gin.Context, err error) {
	lang := getLang(c)

	var appErr *AppError
	if errors.As(err, &appErr) {
		status := codeToStatus(appErr.Code)
		c.JSON(status, ErrorResponse{Error: ErrorBody{
			Code:    appErr.Code,
			Message: i18n.Translate(string(appErr.Code), lang),
		}})
		return
	}

	c.JSON(http.StatusInternalServerError, ErrorResponse{Error: ErrorBody{
		Code:    CodeInternal,
		Message: i18n.Translate(string(CodeInternal), lang),
	}})
}

func codeToStatus(code ErrorCode) int {
	switch code {
	case CodeEmailAlreadyExists, CodeTagAlreadyExists:
		return http.StatusConflict
	case CodeInvalidCredentials, CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeUserWithEmailNotFound, CodeNotFound, CodeTagNotFound, CodeTaskNotFound, CodeProfileNotFound:
		return http.StatusNotFound
	case CodeValidation:
		return http.StatusUnprocessableEntity
	case CodeBadRequest, CodeInvalidID, CodeFormParseFailed, CodeFileTypeNotAllowed:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
