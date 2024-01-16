package usecase

import (
	"errors"
	"interview/domain"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type cartUsecase struct {
	cartRepository     domain.CartRepository
	cartItemRepository domain.CartItemRepository
	contextTimeout     time.Duration
}

func NewCartUsecase(cartRepository domain.CartRepository, cartItemRepository domain.CartItemRepository, timeout time.Duration) domain.CartUsecase {
	return &cartUsecase{
		cartRepository:     cartRepository,
		cartItemRepository: cartItemRepository,
		contextTimeout:     timeout,
	}
}

var itemPriceMapping = map[string]float64{
	"shoe":  100,
	"purse": 200,
	"bag":   300,
	"watch": 300,
}

func (pu *cartUsecase) AddItemToCart(c *gin.Context, cartSessionId string, itemForm domain.CartItemForm) error {
	var isCartNew bool
	var cartEntity domain.CartEntity
	sessionId := cartSessionId
	_, err := pu.cartRepository.FindBySessionId(c, sessionId)

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// c.Redirect(302, "/")
			return err
		}
		isCartNew = true
		cartEntity = domain.CartEntity{
			SessionID: sessionId,
			Status:    domain.CartOpen,
		}
		pu.cartRepository.Create(c, &cartEntity)
	}

	if err != nil {
		return err
	}

	item, ok := itemPriceMapping[itemForm.Product]
	if !ok {
		return errors.New("invalid item name")
	}

	quantity, err := strconv.ParseInt(itemForm.Quantity, 10, 0)
	if err != nil {
		return errors.New("invalid quantity")
	}

	var cartItemEntity domain.CartItemEntity
	if isCartNew {
		cartItemEntity = domain.CartItemEntity{
			CartID:      cartEntity.ID,
			ProductName: itemForm.Product,
			Quantity:    int(quantity),
			Price:       item * float64(quantity),
		}
		pu.cartItemRepository.Create(c, &cartItemEntity)
	} else {
		_, err = pu.cartRepository.FindByProductName(c, cartEntity.ID, itemForm.Product)

		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			cartItemEntity = domain.CartItemEntity{
				CartID:      cartEntity.ID,
				ProductName: itemForm.Product,
				Quantity:    int(quantity),
				Price:       item * float64(quantity),
			}
			pu.cartItemRepository.Create(c, &cartItemEntity)

		} else {
			cartItemEntity.Quantity += int(quantity)
			cartItemEntity.Price += item * float64(quantity)
			pu.cartItemRepository.Save(c, cartItemEntity)
		}
	}

	return nil
}

func (pu *cartUsecase) DeleteCartItem(c *gin.Context, cartId string) error {
	if cartId == "" {
		return errors.New("Cart is Empty")
	}

	cookie, _ := c.Request.Cookie("ice_session_id")

	cartEntity, err := pu.cartRepository.FindBySessionId(c, cookie.Value)
	if err != nil {
		// c.Redirect(302, "/")
		return errors.New("Cart is Incompatible with session value is Empty")
	}

	if cartEntity.Status == domain.CartClosed {
		c.Redirect(302, "/")
		return errors.New("Cart status is CLOSED")
	}

	cartItemID, err := strconv.Atoi(cartId)
	if err != nil {
		// c.Redirect(302, "/")
		return errors.New("Cart Item Value is not INT")
	}

	_, errCartItem := pu.cartItemRepository.FindById(c, cartItemID)
	if errCartItem != nil {
		// c.Redirect(302, "/")
		return errors.New("Cart is Empty")
	}

	pu.cartItemRepository.Delete(c, cartItemID)
	// c.Redirect(302, "/")
	return nil
}

func (pu *cartUsecase) GetCartItemData(c *gin.Context, sessionID string) ([]map[string]interface{}, error) {
	cartEntity, err := pu.cartRepository.FindBySessionId(c, sessionID)
	if err != nil {
		return nil, err
	}

	cartItems, err := pu.cartItemRepository.FindByCartId(c, int(cartEntity.ID))
	if err != nil {
		return nil, err
	}

	items := make([]map[string]interface{}, 1)
	for _, cartItem := range cartItems {
		item := map[string]interface{}{
			"ID":       cartItem.ID,
			"Quantity": cartItem.Quantity,
			"Price":    cartItem.Price,
			"Product":  cartItem.ProductName,
		}

		items = append(items, item)
	}
	return items, nil
}
