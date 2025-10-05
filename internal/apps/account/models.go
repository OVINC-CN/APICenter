package account

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ovinc-cn/apicenter/v2/internal/config"
	"github.com/ovinc-cn/apicenter/v2/pkg/mysql"
	"github.com/ovinc-cn/apicenter/v2/pkg/password"
	"github.com/ovinc-cn/apicenter/v2/pkg/redis"
)

type User struct {
	ID            uint64         `json:"id" gorm:"primaryKey;autoIncrement;"`
	Username      string         `json:"username" gorm:"type:varchar(32);unique;"`
	NickName      sql.NullString `json:"nick_name" gorm:"type:varchar(32);default:null;unique;"`
	Password      string         `json:"password" gorm:"type:varchar(255);"`
	DateJoined    time.Time      `json:"date_joined" gorm:"type:datetime;autoCreateTime;"`
	LastLogin     time.Time      `json:"last_login" gorm:"type:datetime;"`
	PhoneNumber   sql.NullString `json:"phone_number" gorm:"type:varchar(32);default:null;unique;"`
	EmailAddress  sql.NullString `json:"email_address" gorm:"type:varchar(255);default:null;unique;"`
	WeChatOpenID  sql.NullString `json:"wechat_open_id" gorm:"column:wechat_open_id;type:varchar(255);default:null;unique;"`
	WeChatUnionID sql.NullString `json:"wechat_union_id" gorm:"column:wechat_union_id;type:varchar(255);default:null;unique;"`
	Avatar        sql.NullString `json:"avatar" gorm:"type:varchar(255);default:null;"`
}

func (u *User) ExactByUsername(ctx context.Context, username string) error {
	if err := mysql.Select(ctx, mysql.DB(), u, "username = ?", username); err != nil {
		return err
	}
	if u.ID == 0 {
		return fmt.Errorf("user not found: %s", username)
	}
	return nil
}

func (u *User) ExactByUsernameAndPassword(ctx context.Context, username, passWd string) error {
	// load user from db
	if err := u.ExactByUsername(ctx, username); err != nil {
		return err
	}
	// check password
	isMatch, err := password.CheckPassword(passWd, u.Password)
	if err != nil {
		return err
	}
	if !isMatch {
		return fmt.Errorf("password not match")
	}
	// update last login time
	u.LastLogin = time.Now()
	if err := mysql.Update(ctx, mysql.DB(), u, "last_login", time.Now().UTC()); err != nil {
		return err
	}
	return nil
}

func (u *User) MakeNewToken(ctx context.Context) (string, error) {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d", time.Now().UnixMilli())))
	token := fmt.Sprintf("%x#%s", h.Sum(nil), uuid.NewString())
	if _, err := redis.Set(ctx, redis.Client(), u.TokenCacheKey(token), u.Username, time.Duration(u.CacheKeyTimeout())*time.Second); err != nil {
		return "", err
	}
	return token, nil
}

func (u *User) ValidateToken(ctx context.Context, token string) error {
	// load from cache
	username, err := redis.Get(ctx, redis.Client(), u.TokenCacheKey(token))
	if err != nil {
		return err
	}
	if username == "" {
		return fmt.Errorf("invalid token: %s", token)
	}

	// load user from db
	return u.ExactByUsername(ctx, username)
}

func (u *User) TokenCacheKey(token string) string {
	return fmt.Sprintf("account:user:token:%s", token)
}

func (u *User) CacheKeyTimeout() int {
	return config.Config.AppAccount.AuthTokenTimeout
}

func (u *User) CookieName() string {
	return config.Config.AppAccount.AuthTokenCookieName
}

func (u *User) CookieDomain() string {
	return config.Config.API.CookieDomain
}

func (u *User) CookieSecure() bool {
	return config.Config.API.CookieSecure
}

func (u *User) CookieHTTPOnly() bool {
	return config.Config.API.CookieHTTPOnly
}
