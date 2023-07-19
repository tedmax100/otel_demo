package entity

import (
	paymentEntity "otel_demo/external_system/payment/entity"
	"otel_demo/product/entity"
	"time"

	"github.com/gin-gonic/gin"
)

type Order struct {
	BuyerId     uint64
	BuyerName   string
	Products    []entity.Product
	State       uint8
	PaymentInfo paymentEntity.Payment
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (o *Order) CreateOrder(c *gin.Context) {
	//
}

func (o *Order) PayOrder(c *gin.Context) {
	//
}
