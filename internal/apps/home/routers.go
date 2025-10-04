package home

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/pkg/apiMixin"
	"github.com/ovinc-cn/apicenter/v2/pkg/mysql"
	"github.com/ovinc-cn/apicenter/v2/pkg/redis"
)

func HealthZ(c *gin.Context) {
	// trace
	ctx := c.MustGet(apiMixin.TraceContextKey).(context.Context)

	// check for db
	if mysql.DB().Ping() != nil {
		apiMixin.ResponseError(c, MySQLConnFailed)
		return
	}

	// check for redis
	if err := redis.Ping(ctx, redis.Client()); err != nil {
		apiMixin.ResponseError(c, RedisConnFailed)
		return
	}

	apiMixin.ResponseSuccess(c, nil)
}
