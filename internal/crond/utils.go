package crond

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/ovinc-cn/apicenter/v2/pkg/asyncTask"
	"github.com/ovinc-cn/apicenter/v2/pkg/logger"
	"github.com/ovinc-cn/apicenter/v2/pkg/trace"
	"github.com/ovinc-cn/apicenter/v2/pkg/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
)

func NewTask(ctx context.Context, taskType TaskType, data interface{}, customOps ...asynq.Option) (*asynq.Task, error) {
	// init trace
	ctx, span := trace.StartSpan(ctx, fmt.Sprintf("NewTask#%s", taskType), trace.SpanKindInternal)
	defer span.End()

	// init payload
	payload := asyncTask.TaskPayload{Carrier: propagation.MapCarrier{}, Data: data}

	// inject carrier
	otel.GetTextMapPropagator().Inject(ctx, payload.Carrier)

	// build payload
	payloadData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// build options
	ops := []asynq.Option{
		asynq.Queue(getQueue(taskType)),
		asynq.MaxRetry(0),
	}
	ops = append(ops, getOps(taskType)...)
	ops = append(ops, customOps...)

	// build task
	return asynq.NewTask(taskType, payloadData, ops...), nil
}

func RunAsyncTask(ctx context.Context, taskType TaskType, data interface{}, ops ...asynq.Option) error {
	// trace
	ctx, span := trace.StartSpan(ctx, fmt.Sprintf("RunAsyncTask#%s", taskType), trace.SpanKindProducer)
	defer span.End()

	// init task
	task, err := NewTask(ctx, taskType, data, ops...)
	if err != nil {
		return err
	}

	// run task
	taskInfo, err := AsyncClient.Enqueue(task)
	if err != nil {
		if errors.Is(err, asynq.ErrDuplicateTask) {
			logger.Logger(ctx).Warn(fmt.Sprintf("[RunAsyncTask] duplicate %s task", taskType))
			return nil
		}
		span.SetStatus(codes.Error, err.Error())
		logger.Logger(ctx).Error(fmt.Sprintf("[RunAsyncTask] enqueue %s task failed\nerr: %s", taskType, err))
		return err
	}
	logger.Logger(ctx).Info(fmt.Sprintf("[RunAsyncTask] enqueue %s task success\ntask: %s", taskType, utils.ForceStringMaxLength(taskInfo)))
	return nil
}
