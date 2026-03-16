package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wasilisk/doit-api/internal/i18n"
)

func Lang() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Accept-Language")
		lang := parseAcceptLanguage(header)
		c.Set("lang", lang)
		c.Next()
	}
}

func parseAcceptLanguage(header string) i18n.Lang {
	if header == "" {
		return i18n.EN
	}

	primary := strings.SplitN(header, ",", 2)[0]
	primary = strings.SplitN(primary, ";", 2)[0]
	primary = strings.TrimSpace(primary)
	tag := strings.SplitN(primary, "-", 2)[0]
	tag = strings.ToLower(tag)

	return i18n.ParseLang(tag)
}
