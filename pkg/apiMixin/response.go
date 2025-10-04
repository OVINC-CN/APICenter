package apiMixin

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/pkg/trace"
)

func Response(c *gin.Context, status int, msg string, data interface{}) {
	// load context
	ctx := c.MustGet(TraceContextKey).(context.Context)
	span := trace.SpanFromContext(ctx)
	traceID := span.SpanContext().TraceID().String()

	// response
	c.JSON(status, gin.H{"trace_id": traceID, "message": msg, "data": data})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	Response(c, http.StatusOK, "", data)
}

func ResponseError(c *gin.Context, error *APIError) {
	Response(c, error.Code, error.Msg, nil)
}
