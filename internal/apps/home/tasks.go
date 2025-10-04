package home

import (
	"context"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/ovinc-cn/apicenter/v2/pkg/redis"
	"github.com/ovinc-cn/apicenter/v2/pkg/trace"
	"go.opentelemetry.io/otel/codes"
)

var DebugTaskLockKey = "task:home:debug"
var DebugTaskLockTTL = 1 * time.Hour

func DebugTask(ctx context.Context, t *asynq.Task) error {
	// trace
	ctx, span := trace.StartSpan(ctx, fmt.Sprintf("Task#%s", t.Type()), trace.SpanKindConsumer)
	defer span.End()

	// lock
	lock := redis.NewLock(ctx, redis.Client(), DebugTaskLockKey, DebugTaskLockTTL)
	defer lock.Release()
	if err := lock.Lock(); err != nil {
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	return nil
}
