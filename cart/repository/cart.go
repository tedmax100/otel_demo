package repository

import (
	"context"
	"otel_demo/cart/entity"

	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewCartRepository),
)

type CartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) AddItemToCart(ctx context.Context, cart entity.Cart) error {
	_, err := r.db.ExecContext(ctx,
		"REPLACE INTO cart(`buyer_id`, `buyer_name`,`items`, `total_price`, `updated_at`) VALUES ($1,$2,$3,$4,$5)")
	return err
}

func (r *CartRepository) GetCart(ctx context.Context, buyerId uint64) (entity.Cart, error) {
	cart := entity.Cart{}
	err := r.db.GetContext(ctx, &cart, "SELECT * FROM cart WHERE buyer_id=$1", buyerId)
	return cart, err
}
