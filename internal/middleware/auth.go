package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	apperror "github.com/wasilisk/doit-api/internal/app_error"
	"github.com/wasilisk/doit-api/internal/utils"
)

func Auth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			apperror.HandleError(c, apperror.New(apperror.CodeUnauthorized))
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			apperror.HandleError(c, apperror.New(apperror.CodeUnauthorized))
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1], jwtSecret)
		if err != nil {
			apperror.HandleError(c, apperror.New(apperror.CodeUnauthorized))
			c.Abort()
			return
		}

		userID, err := uuid.Parse(claims.UserID)
		if err != nil {
			apperror.HandleError(c, apperror.New(apperror.CodeBadRequest))
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
