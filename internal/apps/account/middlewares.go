package account

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/pkg/apiMixin"
	"github.com/ovinc-cn/apicenter/v2/pkg/logger"
	"github.com/ovinc-cn/apicenter/v2/pkg/trace"
	"go.opentelemetry.io/otel/attribute"
)

func LoginRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// init
		ctx, span := trace.StartSpan(apiMixin.GetTraceCtx(c), "LoginRequiredMiddleware", trace.SpanKindInternal)
		defer span.End()

		// user
		user := &User{}

		// get cookie
		cookie, err := c.Cookie(user.CookieName())
		if err != nil || cookie == "" {
			apiMixin.ResponseUnauthorized(c)
			return
		}

		// validate token
		if err = user.ValidateToken(ctx, cookie); err != nil {
			logger.Logger(ctx).Info(fmt.Sprintf("[LoginRequiredMiddleware] validate token failed\n%v", err))
			apiMixin.ResponseUnauthorized(c)
			return
		}

		// set user to context
		apiMixin.SetUsername(c, user.Username)
		apiMixin.SetUserID(c, user.ID)
		if err := apiMixin.SetUser(c, user); err != nil {
			logger.Logger(ctx).Error(fmt.Sprintf("[LoginRequiredMiddleware] set user to context failed\n%v", err))
			apiMixin.ResponseInternalError(c)
			return
		}

		// trace
		span.SetAttributes(
			attribute.String("username", user.Username),
			attribute.String("user_id", fmt.Sprintf("%d", user.ID)),
		)
		span.End()

		// next
		c.Next()
	}
}
