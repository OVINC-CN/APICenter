package trace

import otTrace "go.opentelemetry.io/otel/trace"

const (
	AttributeRequestURI           = "request.uri"
	AttributeRequestRemoteAddr    = "request.remote_addr"
	AttributeRequestContentLength = "request.content_length"
	AttributeStatusCode           = "status.code"
	AttributeDBInstance           = "db.instance"
	AttributeDBIP                 = "db.ip"
	AttributeDBSystem             = "db.system"
)

const (
	SpanKindInternal = otTrace.SpanKindInternal
	SpanKindServer   = otTrace.SpanKindServer
	SpanKindClient   = otTrace.SpanKindClient
	SpanKindProducer = otTrace.SpanKindProducer
	SpanKindConsumer = otTrace.SpanKindConsumer
)
