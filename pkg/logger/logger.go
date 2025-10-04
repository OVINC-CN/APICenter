package logger

import (
	"context"
	"log"

	"github.com/ovinc-cn/apicenter/v2/pkg/cfg"
	"github.com/ovinc-cn/apicenter/v2/pkg/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	var err error
	config := zap.Config{
		Level:       LogLevel(),
		Development: cfg.AppDebug(),
		Encoding:    cfg.AppLogEncoding(),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "logger",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		},
		OutputPaths: []string{cfg.AppLogOutput()},
	}
	if logger, err = config.Build(); err != nil {
		log.Fatalf("[Logger] init failed; %s", err)
	}
}

func Logger(ctx context.Context) *zap.Logger {
	span := trace.SpanFromContext(ctx)
	spanCtx := span.SpanContext()
	return logger.With(
		zap.String("trace_id", spanCtx.TraceID().String()),
		zap.String("span_id", spanCtx.SpanID().String()),
	)
}
