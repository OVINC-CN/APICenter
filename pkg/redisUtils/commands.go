package redisUtils

import (
	"context"
	"errors"
	"time"

	"github.com/ovinc-cn/apicenter/v2/pkg/commonUtils"
	"github.com/ovinc-cn/apicenter/v2/pkg/configUtils"
	"github.com/ovinc-cn/apicenter/v2/pkg/traceUtils"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func buildAttributes(attrs ...attribute.KeyValue) []attribute.KeyValue {
	return append(
		[]attribute.KeyValue{
			attribute.String(traceUtils.AttributeDBSystem, "redis"),
			attribute.String(traceUtils.AttributeDBIP, configUtils.RedisAddr()),
			attribute.Int(traceUtils.AttributeDBInstance, configUtils.RedisDB()),
		},
		attrs...,
	)
}

func Set(ctx context.Context, conn redis.Cmdable, k string, v interface{}, ex time.Duration) (string, error) {
	// trace
	ctx, span := traceUtils.StartSpan(ctx, "Redis#Set", trace.SpanKindClient)
	defer span.End()

	// add data
	span.SetAttributes(
		buildAttributes(
			attribute.String("key", k),
			attribute.String("val", commonUtils.ForceStringMaxLength(v)),
			attribute.String("ex", ex.String()),
		)...,
	)

	// exec
	cmd := conn.Set(ctx, k, v, ex)
	if err := cmd.Err(); err != nil {
		span.SetStatus(codes.Error, err.Error())
		return "", err
	}
	return cmd.Val(), nil
}

func SetNX(ctx context.Context, conn redis.Cmdable, k string, v interface{}, ex time.Duration) (bool, error) {
	// trace
	ctx, span := traceUtils.StartSpan(ctx, "Redis#SetNX", trace.SpanKindClient)
	defer span.End()

	// add data
	span.SetAttributes(
		buildAttributes(
			attribute.String("key", k),
			attribute.String("val", commonUtils.ForceStringMaxLength(v)),
			attribute.String("ex", ex.String()),
		)...,
	)

	// exec
	cmd := conn.SetNX(ctx, k, v, ex)
	if err := cmd.Err(); err != nil {
		span.SetStatus(codes.Error, err.Error())
		return false, err
	}
	return cmd.Val(), nil
}

func Get(ctx context.Context, conn redis.Cmdable, k string) (string, error) {
	// trace
	ctx, span := traceUtils.StartSpan(ctx, "Redis#Get", trace.SpanKindClient)
	defer span.End()

	// add data
	span.SetAttributes(buildAttributes(attribute.String("key", k))...)

	// exec
	cmd := conn.Get(ctx, k)
	if err := cmd.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		span.SetStatus(codes.Error, err.Error())
		return "", err
	}
	// result
	return cmd.Val(), nil
}

func Del(ctx context.Context, conn redis.Cmdable, keys ...string) (int64, error) {
	// trace
	ctx, span := traceUtils.StartSpan(ctx, "Redis#Del", trace.SpanKindClient)
	defer span.End()

	// add data
	span.SetAttributes(buildAttributes(attribute.StringSlice("keys", keys))...)

	// exec
	cmd := conn.Del(ctx, keys...)
	if err := cmd.Err(); err != nil {
		span.SetStatus(codes.Error, err.Error())
		return 0, err
	}
	// result
	return cmd.Val(), nil
}
