package apiMixin

import (
	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/pkg/i18n"
)

func Translate(c *gin.Context, messageID string, data map[string]string) string {
	// load language
	lang, _ := c.Cookie(I18nLangCookieKey)
	if lang == "" {
		lang = i18n.DefaultLang.String()
	}

	// translate
	return i18n.Translate(lang, messageID, data)
}
