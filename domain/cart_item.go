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
	AddItemToCart(c *gin.Context, sessionId string)
}

type CartItemRepository interface {
	Create(c context.Context, cart *CartItemEntity) error
	FindById(c context.Context, id int) (CartItemEntity, error)
	FindByCartId(c context.Context, id int) ([]CartItemEntity, error)
	Save(c context.Context, cartItemEntity CartItemEntity) error
	Delete(c context.Context, id int) error
}

type CartItemForm struct {
	Product  string `form:"product"   binding:"required"`
	Quantity string `form:"quantity"  binding:"required"`
}
