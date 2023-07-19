package entity

import "otel_demo/product/entity"

type CartItem struct {
	Product  entity.Product
	Quantity uint
}

type Cart struct {
	BuyerId    uint64
	BuyerName  string
	Items      []CartItem
	TotalPrice uint
}

func (cart *Cart) AddItemToCart(item CartItem) {
	cart.Items = append(cart.Items, item)
	cart.calculateCartSummary(item)
}

func (cart *Cart) calculateCartSummary(item CartItem) {
	cart.TotalPrice += item.Product.Price * item.Quantity
}
