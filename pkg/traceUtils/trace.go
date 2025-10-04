package traceUtils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ovinc-cn/apicenter/v2/pkg/configUtils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer
var traceProvider *sdktrace.TracerProvider

func init() {
	// init ctx
	ctx := context.Background()

	// load process name
	if len(os.Args) < 2 {
		log.Fatalf("[Trace] process name not found")
	}
	processName := os.Args[1]

	// resource
	r, err := resource.New(ctx, resource.WithAttributes(
		semconv.ServiceName(fmt.Sprintf("%s-%s", configUtils.AppName(), processName))),
	)
	if err != nil {
		log.Fatalf("[Trace] create resource failed; %s", err)
	}

	// init exporter
	e, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint(configUtils.TraceEndpoint()), otlptracegrpc.WithInsecure())
	if err != nil {
		log.Fatalf("[Trace] create otlp grpc exporter failed; %s", err)
	}

	// init trace provider
	traceProvider = sdktrace.NewTracerProvider(sdktrace.WithBatcher(e), sdktrace.WithResource(r))
	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// init tracer
	tracer = otel.Tracer(configUtils.AppName())
}

func OnExit() {
	_ = traceProvider.Shutdown(context.Background())
}
