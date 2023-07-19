package repository

import (
	"otel_demo/order/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewProductRepository),
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) SubmitOrder(order entity.Order) {

}

func (r *ProductRepository) OrderInventoryValidated(order entity.Order) {

}

func (r *ProductRepository) PayOrder(order entity.Order) {

}

func (r *ProductRepository) OrderPayCompleted(order entity.Order) {

}

func (r *ProductRepository) OrderNotified(order entity.Order) {

}
