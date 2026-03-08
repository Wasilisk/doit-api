package handlerutils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Bind[T any](c *gin.Context) (*T, bool) {
	var req T
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, false
	}
	return &req, true
}
