package asyncTask

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/hibiken/asynq"
	"github.com/ovinc-cn/apicenter/v2/pkg/logger"
	"github.com/ovinc-cn/apicenter/v2/pkg/trace"
	"github.com/ovinc-cn/apicenter/v2/pkg/utils"
	"go.opentelemetry.io/otel"
)

type ScheduleTaskFailed struct {
	TaskName string
}

func (e ScheduleTaskFailed) Error() string {
	return fmt.Sprintf("%sFailed", e.TaskName)
}

func ObserveMiddleware(next asynq.Handler) asynq.Handler {
	propagator := otel.GetTextMapPropagator()
	return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
		// parse payload to map
		var payload TaskPayload
		if err := json.Unmarshal(task.Payload(), &payload); err != nil {
			logger.Logger(ctx).Error(
				fmt.Sprintf("[%s] failed\npayload: %s\nerr: %s", task.Type(), utils.ForceStringMaxLength(task.Payload()), err),
			)
			return err
		}

		// load trace if exists
		if len(payload.Carrier) > 0 {
			ctx = propagator.Extract(ctx, payload.Carrier)
		}

		// span
		ctx, span := trace.StartSpan(ctx, fmt.Sprintf("ObserveTask#%s", task.Type()), trace.SpanKindInternal)
		defer span.End()

		// pre exec
		logger.Logger(ctx).Info(fmt.Sprintf("[%s] start", task.Type()))

		// execute
		err := next.ProcessTask(ctx, task)

		// post exec
		if err != nil {
			logger.Logger(ctx).Error(
				fmt.Sprintf("[%s] failed\npayload: %s\nerr: %s", task.Type(), utils.ForceStringMaxLength(task.Payload()), err),
			)
		} else {
			logger.Logger(ctx).Info(fmt.Sprintf("[%s] finished", task.Type()))
		}
		return err
	})
}

func RecoverMiddleware(next asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) (err error) {
		defer func() {
			if r := recover(); r != nil {
				ctx, span := trace.StartSpan(ctx, fmt.Sprintf("Task#%s#Panic", task.Type()), trace.SpanKindInternal)
				defer span.End()

				var stack []byte
				var length int
				stack = make([]byte, 4096)
				length = runtime.Stack(stack, true)
				stack = stack[:length]

				logger.Logger(ctx).Error(
					fmt.Sprintf("[%s] panic\nerr: %s\n%s", task.Type(), utils.ForceStringMaxLength(r), stack),
				)
			}
		}()
		return next.ProcessTask(ctx, task)
	})
}
