package handlerutils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	apperror "github.com/wasilisk/doit-api/internal/app_error"
)

func tagToCode(tag string) apperror.ErrorCode {
	switch validationTag(tag) {
	case tagRequired:
		return "FIELD_REQUIRED"
	case tagEmail:
		return "FIELD_INVALID_EMAIL"
	case tagMin:
		return "FIELD_MIN"
	case tagMax:
		return "FIELD_MAX"
	default:
		return apperror.ErrorCode(fmt.Sprintf("FIELD_INVALID_%s", strings.ToUpper(tag)))
	}
}

func tagToContext(tag, param string) map[string]string {
	switch validationTag(tag) {
	case tagMin:
		return map[string]string{"min": param}
	case tagMax:
		return map[string]string{"max": param}
	default:
		return nil
	}
}

func tagToMessage(tag, field, param string) string {
	switch validationTag(tag) {
	case tagRequired:
		return fmt.Sprintf("%s is required", toSnakeCase(field))
	case tagEmail:
		return "Please enter a valid email address"
	case tagMin:
		return fmt.Sprintf("%s must be at least %s characters", toSnakeCase(field), param)
	case tagMax:
		return fmt.Sprintf("%s must be at most %s characters", toSnakeCase(field), param)
	default:
		return fmt.Sprintf("%s is invalid", toSnakeCase(field))
	}
}

func toFieldErrors(err error) []apperror.FieldError {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return nil
	}

	fields := make([]apperror.FieldError, len(ve))
	for i, e := range ve {
		fields[i] = apperror.FieldError{
			Field:   toSnakeCase(e.Field()),
			Code:    tagToCode(e.Tag()),
			Message: tagToMessage(e.Tag(), e.Field(), e.Param()),
			Context: tagToContext(e.Tag(), e.Param()),
		}
	}
	return fields
}

func toSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}
