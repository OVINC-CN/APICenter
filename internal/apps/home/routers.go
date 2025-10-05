package home

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/pkg/apiMixin"
	"github.com/ovinc-cn/apicenter/v2/pkg/logger"
	"github.com/ovinc-cn/apicenter/v2/pkg/mysql"
	"github.com/ovinc-cn/apicenter/v2/pkg/redis"
)

// HealthZ godoc
//
//	@Tags		Home
//	@Produce	json
//	@Success	200	{object}	apiMixin.ResponseModel
//	@Router		/v1/healthz [get]
func HealthZ(c *gin.Context) {
	// trace
	ctx := apiMixin.GetTraceCtx(c)

	// check for db
	sqlDB, err := mysql.DB().DB()
	if err != nil {
		logger.Logger(ctx).Error(fmt.Sprintf("[HealthZ] get sql db failed\n%s", err.Error()))
		apiMixin.ResponseInternalError(c)
		return
	}
	if err := sqlDB.Ping(); err != nil {
		logger.Logger(ctx).Error(fmt.Sprintf("[HealthZ] ping mysql failed\n%s", err.Error()))
		apiMixin.ResponseInternalError(c)
		return
	}

	// check for redis
	if err := redis.Ping(ctx, redis.Client()); err != nil {
		logger.Logger(ctx).Error(fmt.Sprintf("[HealthZ] ping redis failed\n%s", err.Error()))
		apiMixin.ResponseInternalError(c)
		return
	}

	apiMixin.ResponseSuccess(c, nil)
}
