package domain

import (
	"context"

	"gorm.io/gorm"
)

const (
	CartOpen   = "open"
	CartClosed = "closed"
)

type CartEntity struct {
	gorm.Model
	Total     float64
	SessionID string
	Status    string
}

type CartUsecase interface {
	GetCartData(c context.Context) CartEntity
	AddItemToCart(c context.Context, item string, card string) ([]CartEntity, error)
	DeleteCartItem(c context.Context, cartItemID string) error
}

type CartRepository interface {
	Create(c context.Context, cart *CartEntity) error
	Where(c context.Context, criteria string) []CartEntity
	Delete(c context.Context, cart *CartEntity) error
}
