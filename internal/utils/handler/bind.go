package handlerutils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperror "github.com/wasilisk/doit-api/internal/app_error"
	"github.com/wasilisk/doit-api/internal/i18n"
)

func Bind[T any](c *gin.Context) (*T, bool) {
	var req T
	return bindWith(c, &req, func() error {
		return c.ShouldBind(&req)
	})
}

func BindJSON[T any](c *gin.Context) (*T, bool) {
	var req T
	return bindWith(c, &req, func() error {
		return c.ShouldBindJSON(&req)
	})
}

func bindWith[T any](c *gin.Context, req *T, bindFn func() error) (*T, bool) {
	if err := bindFn(); err != nil {
		handleBindError(c, err)
		return nil, false
	}
	return req, true
}

func handleBindError(c *gin.Context, err error) {
	lang := getLang(c)

	if fields := toFieldErrors(err, lang); fields != nil {
		c.JSON(http.StatusUnprocessableEntity, apperror.ErrorResponse{
			Error: apperror.ErrorBody{
				Code:    apperror.CodeValidation,
				Message: i18n.Translate(string(apperror.CodeValidation), lang),
				Fields:  fields,
			},
		})
		return
	}
	c.JSON(http.StatusBadRequest, apperror.ErrorResponse{
		Error: apperror.ErrorBody{
			Code:    apperror.CodeBadRequest,
			Message: i18n.Translate(string(apperror.CodeBadRequest), lang),
		},
	})
}

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
