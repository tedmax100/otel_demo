package email

import (
	"context"
	"otel_demo/order/entity"

	"go.opentelemetry.io/otel/trace"
)

type EMail struct {
	Sender   string
	Receiver string
	Body     string
}

func (e EMail) EMailOrder(ctx context.Context, orer entity.Order) error {
	span := trace.SpanFromContext(ctx)
	defer span.End()
	return nil
}
