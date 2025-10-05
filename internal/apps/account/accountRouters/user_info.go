package accountRouters

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/internal/apps/account"
	"github.com/ovinc-cn/apicenter/v2/pkg/apiMixin"
	"github.com/ovinc-cn/apicenter/v2/pkg/cfg"
)

type UserInfoData struct {
	UserID    uint64 `json:"user_id"`
	Username  string `json:"username"`
	NickName  string `json:"nick_name"`
	LastLogin string `json:"last_login"`
}

type UserInfoResponse struct {
	apiMixin.ResponseModel
	Data UserInfoData `json:"data"`
}

// UserInfo godoc
//
//	@Tags		Account
//	@Produce	json
//	@Success	200	{object}	UserInfoResponse
//	@Router		/v1/account/user-info [get]
func UserInfo(c *gin.Context) {
	// load user from context
	var user account.User
	if err := apiMixin.GetUser(c, &user); err != nil {
		apiMixin.ResponseInternalError(c)
		return
	}

	// response
	apiMixin.ResponseSuccess(
		c,
		UserInfoData{
			UserID:    user.ID,
			Username:  user.Username,
			NickName:  user.GetNickName(),
			LastLogin: user.LastLogin.In(cfg.AppTimezone()).Format(time.RFC3339),
		},
	)
}
