package domain

import (
	"gorm.io/gorm"
)

type CartItemEntity struct {
	gorm.Model
	CartID      uint
	ProductName string
	Quantity    int
	Price       float64
}

type CartItemUsecase interface {
	// GetCartData(c context.Context) CartEntity
	// AddItemToCart(c context.Context, item string, card string) ([]CartEntity, error)
	// DeleteCartItem(c context.Context, cartItemID string) error
}

type CartItemRepository interface {
	// Create(c context.Context, cart *CartEntity) error
	// Where(c context.Context, criteria string) []CartEntity
	// Delete(c context.Context, cart *CartEntity) error
}
