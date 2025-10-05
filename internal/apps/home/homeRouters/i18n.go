package homeRouters

import (
	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/internal/config"
	"github.com/ovinc-cn/apicenter/v2/pkg/apiMixin"
)

type ChangeLangRequest struct {
	Lang string `json:"lang" binding:"required,oneof=zh-Hans en"`
}

// ChangeLang godoc
//
//	@Tags		Home
//	@Param		request	body	ChangeLangRequest	true "request body"
//	@Produce	json
//	@Success	200	{object}	apiMixin.ResponseModel
//	@Router		/v1/language [put]
func ChangeLang(c *gin.Context) {
	// req
	var req ChangeLangRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiMixin.ResponseBadRequest(c, err.Error())
		return
	}

	// set cookie
	c.SetCookie(
		apiMixin.I18nLangCookieKey,
		req.Lang,
		apiMixin.I18nLangCookieTTL,
		"",
		config.Config.API.CookieDomain,
		config.Config.API.CookieSecure,
		config.Config.API.CookieHTTPOnly,
	)
	apiMixin.ResponseSuccess(c, nil)
}
