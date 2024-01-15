package domain

import (
	"context"

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
	// GetCartData(c context.Context) CartItemEntity
	// AddItemToCart(c context.Context, item string, card string) ([]CartItemEntity, error)
	// DeleteCartItem(c context.Context, cartItemID string) error
}

type CartItemRepository interface {
	Create(c context.Context, cart *CartItemEntity) error
	Where(c context.Context, criteria string) []CartItemEntity
	Save(c context.Context, CartItemEntity string) error
	Delete(c context.Context, cart *CartItemEntity) error
}
