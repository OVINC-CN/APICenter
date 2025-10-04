package crond

import "go.opentelemetry.io/otel/propagation"

type TaskPayload struct {
	Carrier propagation.MapCarrier
	Data    interface{}
}
