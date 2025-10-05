package accountRouters

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/internal/apps/account"
	"github.com/ovinc-cn/apicenter/v2/pkg/apiMixin"
	"github.com/ovinc-cn/apicenter/v2/pkg/logger"
	"github.com/ovinc-cn/apicenter/v2/pkg/redis"
)

// SignOut godoc
//
//	@Tags		Account
//	@Produce	json
//	@Success	200	{object}	apiMixin.ResponseModel
//	@Router		/v1/account/sign-out [get]
func SignOut(c *gin.Context) {
	// init
	ctx := apiMixin.GetTraceCtx(c)

	// load cookie
	user := &account.User{}
	cookie, err := c.Cookie(user.CookieName())
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			apiMixin.ResponseSuccess(c, nil)
			return
		}
		logger.Logger(ctx).Error(fmt.Sprintf("[SignOut] load cookie failed\n%s", err.Error()))
		apiMixin.ResponseInternalError(c)
		return
	}

	// delete redis cache
	if _, err := redis.Del(ctx, redis.Client(), user.TokenCacheKey(cookie)); err != nil {
		logger.Logger(ctx).Error(fmt.Sprintf("[SignOut] delete redis cache failed\n%s", err.Error()))
		apiMixin.ResponseInternalError(c)
		return
	}

	// delete cookie
	c.SetCookie(user.CookieName(), "", -1, "", user.CookieDomain(), user.CookieSecure(), user.CookieHTTPOnly())

	// response
	apiMixin.ResponseSuccess(c, nil)
}
