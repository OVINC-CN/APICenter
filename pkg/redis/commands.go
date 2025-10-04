package redis

import (
	"context"
	"errors"
	"time"

	"github.com/ovinc-cn/apicenter/v2/pkg/cfg"
	"github.com/ovinc-cn/apicenter/v2/pkg/trace"
	"github.com/ovinc-cn/apicenter/v2/pkg/utils"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func buildAttributes(attrs ...attribute.KeyValue) []attribute.KeyValue {
	return append(
		[]attribute.KeyValue{
			attribute.String(trace.AttributeDBSystem, "redis"),
			attribute.String(trace.AttributeDBIP, cfg.RedisAddr()),
			attribute.Int(trace.AttributeDBInstance, cfg.RedisDB()),
		},
		attrs...,
	)
}

func Ping(ctx context.Context, conn redis.Cmdable) error {
	// trace
	ctx, span := trace.StartSpan(ctx, "Redis#Ping", trace.SpanKindClient)
	defer span.End()

	// exec
	cmd := conn.Ping(ctx)
	if err := cmd.Err(); err != nil {
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return nil
}

func Set(ctx context.Context, conn redis.Cmdable, k string, v interface{}, ex time.Duration) (string, error) {
	// trace
	ctx, span := trace.StartSpan(ctx, "Redis#Set", trace.SpanKindClient)
	defer span.End()

	// add data
	span.SetAttributes(
		buildAttributes(
			attribute.String("key", k),
			attribute.String("val", utils.ForceStringMaxLength(v)),
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
	ctx, span := trace.StartSpan(ctx, "Redis#SetNX", trace.SpanKindClient)
	defer span.End()

	// add data
	span.SetAttributes(
		buildAttributes(
			attribute.String("key", k),
			attribute.String("val", utils.ForceStringMaxLength(v)),
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
	ctx, span := trace.StartSpan(ctx, "Redis#Get", trace.SpanKindClient)
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
	ctx, span := trace.StartSpan(ctx, "Redis#Del", trace.SpanKindClient)
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
