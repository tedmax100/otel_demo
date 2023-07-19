package order

import (
	"otel_demo/order/repository"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewOrderController),
)

type OrderController struct {
	repo repository.OrderRepository
}

func NewOrderController(repo repository.OrderRepository) *OrderController {
	return &OrderController{repo: repo}
}

func (oc *OrderController) Submit(c *gin.Context) {
	ctx := c.Request.Context()
	span := trace.SpanFromContext(otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(c.Request.Header)))
	defer span.End()

}

func (oc *OrderController) PayOrder(c *gin.Context) {
	ctx := c.Request.Context()
	span := trace.SpanFromContext(otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(c.Request.Header)))
	defer span.End()
}
