package domain

import (
	"context"

	"github.com/gin-gonic/gin"
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
	// GetCartData(c *gin.Context) CartItemEntity
	AddItemToCart(c *gin.Context, sessionId string)
	// DeleteCartItem(c *gin.Context, cartItemID string) error
}

type CartItemRepository interface {
	Create(c context.Context, cart *CartItemEntity) error
	Where(c context.Context, criteria string) []CartItemEntity
	Save(c context.Context, CartItemEntity string) error
	Delete(c context.Context, cart *CartItemEntity) error
}
