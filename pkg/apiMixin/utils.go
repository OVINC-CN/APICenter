package apiMixin

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func SetTraceCtx(c *gin.Context, ctx context.Context) {
	c.Set(TraceContextKey, ctx)
}

func GetTraceCtx(c *gin.Context) context.Context {
	return c.MustGet(TraceContextKey).(context.Context)
}

func SetUsername(c *gin.Context, username string) {
	c.Set(UsernameContextKey, username)
}

func GetUsername(c *gin.Context) string {
	return c.MustGet(UsernameContextKey).(string)
}

func SetUserID(c *gin.Context, userID uint64) {
	c.Set(UserIDContextKey, userID)
}

func GetUserID(c *gin.Context) uint64 {
	return c.MustGet(UserIDContextKey).(uint64)
}

func SetUser(c *gin.Context, user interface{}) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	c.Set(UserContextKey, userBytes)
	return nil
}

func GetUser[T interface{}](c *gin.Context, user *T) error {
	userBytes := c.MustGet(UserContextKey).([]byte)
	return json.Unmarshal(userBytes, user)
}
