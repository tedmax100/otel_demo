package payment

import (
	"context"
	entity "otel_demo/external_system/payment/entity"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewPaymentService),
)

type PaymentService struct {
}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (p *PaymentService) PayOrder(ctx context.Context, paymentInfo entity.Payment) (entity.Payment, error) {
	span := trace.SpanFromContext(ctx)
	defer span.End()
	return entity.Payment{}, nil
}
