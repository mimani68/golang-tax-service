package domain

import (
	"context"

	"github.com/gin-gonic/gin"
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
	GetCartData(c *gin.Context) CartEntity
	AddItemToCart(c *gin.Context, item string, card string) (CartEntity, error)
	DeleteCartItem(c *gin.Context, cartItemID string) error
}

type CartRepository interface {
	Create(c context.Context, cart *CartEntity) error
	FindBySessionId(c context.Context, session_id string) (CartEntity, error)
	FindByProductName(c context.Context, cart_id string, product_name string) (CartItemEntity, error)
	Delete(c context.Context, cart *CartEntity) error
}
