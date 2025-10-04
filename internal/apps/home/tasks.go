package home

import (
	"context"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/ovinc-cn/apicenter/v2/pkg/redisUtils"
	"github.com/ovinc-cn/apicenter/v2/pkg/traceUtils"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var DebugTaskLockKey = "task:home:debug"
var DebugTaskLockTTL = 1 * time.Hour

func DebugTask(ctx context.Context, t *asynq.Task) error {
	// trace
	ctx, span := traceUtils.StartSpan(ctx, fmt.Sprintf("Task#%s", t.Type()), trace.SpanKindConsumer)
	defer span.End()

	// lock
	lock := redisUtils.NewLock(ctx, redisUtils.Client(), DebugTaskLockKey, DebugTaskLockTTL)
	defer lock.Release()
	if err := lock.Lock(); err != nil {
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	return nil
}
