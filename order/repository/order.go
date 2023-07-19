package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewOrderRepository),
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) SubmitOrder() {

}

func (r *OrderRepository) OrderInventoryValidated() {

}

func (r *OrderRepository) PayOrder() {

}

func (r *OrderRepository) OrderPayCompleted() {

}

func (r *OrderRepository) OrderNotified() {

}
