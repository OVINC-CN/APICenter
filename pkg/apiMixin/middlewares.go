package apiMixin

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/pkg/logger"
	"github.com/ovinc-cn/apicenter/v2/pkg/trace"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

func ObserveMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// load path
		path := c.FullPath()
		if path == "" {
			path = "Unknown"
		}

		// init trace
		ctx, span := trace.StartSpan(context.Background(), fmt.Sprintf("Router#%s#%s", c.Request.Method, path), trace.SpanKindServer)
		defer span.End()
		span.SetAttributes(
			attribute.String(trace.AttributeRequestURI, c.Request.RequestURI),
			attribute.String(trace.AttributeRequestRemoteAddr, c.Request.RemoteAddr),
			attribute.Int64(trace.AttributeRequestContentLength, c.Request.ContentLength),
		)

		// add trace to context
		SetTraceCtx(c, ctx)

		// record time
		startTime := time.Now()

		// do next
		c.Next()

		// record duration
		duration := time.Since(startTime)

		// response status
		status := c.Writer.Status()
		if status >= http.StatusBadRequest {
			span.SetStatus(codes.Error, fmt.Sprintf("status %d", status))
		}

		// add attributes
		span.SetAttributes(
			semconv.HTTPResponseStatusCode(status),
			attribute.Int(trace.AttributeStatusCode, status),
		)

		// log
		logger.Logger(ctx).Info(
			fmt.Sprintf("[ObserveMiddleware] %d | %v | %s | %s", status, duration, c.Request.Method, c.Request.RequestURI),
		)
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// init
				ctx, span := trace.StartSpan(GetTraceCtx(c), "RecoveryMiddleware", trace.SpanKindInternal)
				defer span.End()

				// set error
				span.SetStatus(codes.Error, fmt.Sprintf("%T", r))

				// log
				logger.Logger(ctx).Error(fmt.Sprintf("[RecoveryMiddleware] server panic\n%v", r))

				// return 500
				ResponseInternalError(c)
			}
		}()
		c.Next()
	}
}
