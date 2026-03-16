package handlerutils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	apperror "github.com/wasilisk/doit-api/internal/app_error"
	"github.com/wasilisk/doit-api/internal/i18n"
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

func tagToMessage(tag, field, param string, lang i18n.Lang) string {
	code := tagToCode(tag)
	template := i18n.Translate(string(code), lang)

	switch validationTag(tag) {
	case tagRequired:
		return fmt.Sprintf(template, toSnakeCase(field))
	case tagEmail:
		return template
	case tagMin:
		return fmt.Sprintf(template, toSnakeCase(field), param)
	case tagMax:
		return fmt.Sprintf(template, toSnakeCase(field), param)
	default:
		fallback := i18n.Translate("FIELD_INVALID", lang)
		return fmt.Sprintf(fallback, toSnakeCase(field))
	}
}

func toFieldErrors(err error, lang i18n.Lang) []apperror.FieldError {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return nil
	}

	fields := make([]apperror.FieldError, len(ve))
	for i, e := range ve {
		fields[i] = apperror.FieldError{
			Field:   toSnakeCase(e.Field()),
			Code:    tagToCode(e.Tag()),
			Message: tagToMessage(e.Tag(), e.Field(), e.Param(), lang),
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
