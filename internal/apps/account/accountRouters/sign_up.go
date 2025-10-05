package accountRouters

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/internal/apps/account"
	"github.com/ovinc-cn/apicenter/v2/internal/config"
	"github.com/ovinc-cn/apicenter/v2/pkg/apiMixin"
	"github.com/ovinc-cn/apicenter/v2/pkg/logger"
	"github.com/ovinc-cn/apicenter/v2/pkg/mysql"
	"github.com/ovinc-cn/apicenter/v2/pkg/password"
)

type SignUpRequest struct {
	SignInRequest
	NickName    string `json:"nickname" binding:"min=4,max=32"`
	PhoneNumber string `json:"phone_number" binding:"required,min=11,max=11"`
}

// SignUp godoc
//
//	@Tags		Account
//	@Param		request	body	SignUpRequest	true "request body"
//	@Produce	json
//	@Success	200	{object}	apiMixin.ResponseModel
//	@Router		/v1/account/sign-up [post]
func SignUp(c *gin.Context) {
	// trace
	ctx := apiMixin.GetTraceCtx(c)

	// req
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apiMixin.ResponseBadRequest(c, err.Error())
		return
	}

	// check username pattern
	if !account.UsernameRegex.MatchString(req.Username) {
		apiMixin.ResponseBadRequest(c, apiMixin.Translate(c, "account.usernameInvalid", nil))
		return
	}

	// check username exists
	var userCount int64
	if err := mysql.Count(
		ctx,
		mysql.DB().Model(&account.User{}).Where("username = ?", req.Username).Or("nick_name = ?", req.NickName),
		&userCount,
	); err != nil {
		logger.Logger(ctx).Error(fmt.Sprintf("[SignUp] check username duplicate failed\n%s", err.Error()))
		apiMixin.ResponseInternalError(c)
		return
	}
	if userCount > 0 {
		apiMixin.ResponseBadRequest(c, apiMixin.Translate(c, "account.usernameDuplicate", nil))
		return
	}

	// check phone number exists
	var phoneCount int64
	if err := mysql.Count(
		ctx,
		mysql.DB().Model(&account.User{}).Where("phone_number = ?", req.PhoneNumber),
		&phoneCount,
	); err != nil {
		logger.Logger(ctx).Error(fmt.Sprintf("[SignUp] check phone number duplicate failed\n%s", err.Error()))
		apiMixin.ResponseInternalError(c)
		return
	}
	if phoneCount > 0 {
		apiMixin.ResponseBadRequest(c, apiMixin.Translate(c, "account.phoneNumberDuplicate", nil))
		return
	}

	// make password
	hashedPassword, err := password.MakePassword(req.Password)
	if err != nil {
		logger.Logger(ctx).Error(fmt.Sprintf("[SignUp] hash password failed\n%s", err.Error()))
		apiMixin.ResponseInternalError(c)
		return
	}

	// insert into db
	user := account.User{
		Username:     req.Username,
		NickName:     sql.NullString{String: req.NickName, Valid: true},
		Password:     hashedPassword,
		DateJoined:   time.Now(),
		LastLogin:    time.Now(),
		PhoneNumber:  sql.NullString{String: req.PhoneNumber, Valid: true},
		EmailAddress: sql.NullString{String: fmt.Sprintf("%s@%s", req.Username, config.Config.AppAccount.EmailDomain), Valid: true},
	}
	if err := mysql.Create(ctx, mysql.DB(), &user); err != nil {
		logger.Logger(ctx).Error(fmt.Sprintf("[SignUp] insert user into db failed\n%s", err.Error()))
		apiMixin.ResponseInternalError(c)
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
