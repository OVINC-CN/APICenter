package trace

import otTrace "go.opentelemetry.io/otel/trace"

const (
	AttributeStatusCode = "status.code"
	AttributeDBInstance = "db.instance"
	AttributeDBIP       = "db.ip"
	AttributeDBSystem   = "db.system"
)

const (
	SpanKindInternal = otTrace.SpanKindInternal
	SpanKindServer   = otTrace.SpanKindServer
	SpanKindClient   = otTrace.SpanKindClient
	SpanKindProducer = otTrace.SpanKindProducer
	SpanKindConsumer = otTrace.SpanKindConsumer
)
