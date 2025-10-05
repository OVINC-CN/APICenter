package mysql

import (
	"context"

	"github.com/ovinc-cn/apicenter/v2/pkg/cfg"
	"github.com/ovinc-cn/apicenter/v2/pkg/trace"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"gorm.io/gorm"
)

func buildAttributes(attributes ...attribute.KeyValue) []attribute.KeyValue {
	return append(
		[]attribute.KeyValue{
			attribute.String(trace.AttributeDBSystem, "mysql"),
			attribute.String(trace.AttributeDBIP, cfg.MySQLAddr()),
			attribute.String(trace.AttributeDBInstance, cfg.MySQLDatabase()),
		},
		attributes...,
	)
}

func Count(ctx context.Context, db *gorm.DB, count *int64) error {
	// trace
	ctx, span := trace.StartSpan(ctx, "MySQL#Count", trace.SpanKindClient)
	defer span.End()

	// exec
	result := db.WithContext(ctx).Count(count)
	err := result.Error
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	}

	// add attributes
	span.SetAttributes(buildAttributes()...)

	return err
}

func Create(ctx context.Context, db *gorm.DB, value interface{}) error {
	// trace
	ctx, span := trace.StartSpan(ctx, "MySQL#Create", trace.SpanKindClient)
	defer span.End()

	// exec
	result := db.WithContext(ctx).Create(value)
	err := result.Error
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	}

	// add attributes
	span.SetAttributes(buildAttributes()...)

	return err
}

func Select(ctx context.Context, db *gorm.DB, dest interface{}, conditions ...interface{}) error {
	// trace
	ctx, span := trace.StartSpan(ctx, "MySQL#Select", trace.SpanKindClient)
	defer span.End()

	// exec
	result := db.WithContext(ctx).Find(dest, conditions...)
	err := result.Error
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	}

	// add attributes
	span.SetAttributes(buildAttributes()...)

	return err
}

func Update[T interface{}](ctx context.Context, db *gorm.DB, model *T, column string, value interface{}) error {
	// trace
	ctx, span := trace.StartSpan(ctx, "MySQL#Update", trace.SpanKindClient)
	defer span.End()

	// exec
	result := db.WithContext(ctx).Model(model).Update(column, value)
	err := result.Error
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	}

	// add attributes
	span.SetAttributes(buildAttributes()...)

	return err
}

func Save(ctx context.Context, db *gorm.DB, value interface{}) error {
	// trace
	ctx, span := trace.StartSpan(ctx, "MySQL#Save", trace.SpanKindClient)
	defer span.End()

	// exec
	result := db.WithContext(ctx).Save(value)
	err := result.Error
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	}

	// add attributes
	span.SetAttributes(buildAttributes()...)

	return err
}
