package handlerutils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperror "github.com/wasilisk/doit-api/internal/app_error"
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
	if fields := toFieldErrors(err); fields != nil {
		c.JSON(http.StatusUnprocessableEntity, apperror.ErrorResponse{
			Error: apperror.ErrorBody{
				Code:    apperror.ErrValidation.Code,
				Message: apperror.ErrValidation.Message,
				Fields:  fields,
			},
		})
		return
	}
	c.JSON(http.StatusBadRequest, apperror.ErrValidation)
}
