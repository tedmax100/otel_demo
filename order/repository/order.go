package repository

import (
	"context"

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

func (r *OrderRepository) SubmitOrder(ctx context.Context) error {
	_, err :=	r.db.ExecContext(ctx, `INSERT INTO order (buyer_id, buyer_name, products, state, payment_info, created_at, updated_at)
	`)
	return err
}

func (r *OrderRepository) OrderInventoryValidated() {

}

func (r *OrderRepository) PayOrder() {

}

func (r *OrderRepository) OrderPayCompleted() {

}

func (r *OrderRepository) OrderNotified() {

}
