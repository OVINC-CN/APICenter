package apiMixin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ovinc-cn/apicenter/v2/pkg/trace"
)

func Response(c *gin.Context, status int, msg string, data interface{}) {
	// load context
	ctx := GetTraceCtx(c)
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

func ResponseBadRequest(c *gin.Context, msg string) {
	Response(c, http.StatusBadRequest, msg, nil)
}

func ResponseUnauthorized(c *gin.Context) {
	Response(c, http.StatusUnauthorized, Translate(c, "apiMixin.unAuthorized", nil), nil)
}

func ResponseForbidden(c *gin.Context) {
	Response(c, http.StatusForbidden, Translate(c, "apiMixin.noPermission", nil), nil)
}

func ResponseNotFound(c *gin.Context) {
	Response(c, http.StatusNotFound, Translate(c, "apiMixin.notFound", nil), nil)
}

func ResponseInternalError(c *gin.Context) {
	Response(c, http.StatusInternalServerError, Translate(c, "apiMixin.internalServerError", nil), nil)
}
