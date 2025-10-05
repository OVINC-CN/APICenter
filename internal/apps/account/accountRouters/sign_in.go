package accountRouters

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/internal/apps/account"
	"github.com/ovinc-cn/apicenter/v2/pkg/apiMixin"
	"github.com/ovinc-cn/apicenter/v2/pkg/logger"
)

type SignInRequest struct {
	Username string `json:"username" binding:"required,min=4,max=32"`
	Password string `json:"password" binding:"required,min=64,max=64"`
}

// SignIn godoc
//
//	@Tags		Account
//	@Param		request	body	SignInRequest	true "request body"
//	@Produce	json
//	@Success	200	{object}	apiMixin.ResponseModel
//	@Router		/v1/account/sign-in [post]
func SignIn(c *gin.Context) {
	// init
	ctx := apiMixin.GetTraceCtx(c)

	// req
	var req SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiMixin.ResponseBadRequest(c, err.Error())
		return
	}

	// load user from db
	var user account.User
	if err := user.ExactByUsernameAndPassword(ctx, req.Username, req.Password); err != nil {
		logger.Logger(ctx).Error(fmt.Sprintf("[SignIn] load user from db failed\n%s", err))
		apiMixin.ResponseBadRequest(c, apiMixin.Translate(c, "account.loginFailed", nil))
		return
	}

	// set up cookie
	token, err := user.MakeNewToken(ctx)
	if err != nil {
		logger.Logger(ctx).Error(fmt.Sprintf("[SignUp] make token failed\n%s", err.Error()))
		apiMixin.ResponseInternalError(c)
		return
	}
	c.SetCookie(user.CookieName(), token, user.CacheKeyTimeout(), "", user.CookieDomain(), user.CookieSecure(), user.CookieHTTPOnly())

	// response
	apiMixin.ResponseSuccess(c, nil)
	return
}
