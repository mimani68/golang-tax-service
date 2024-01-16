package usecase

import (
	"errors"
	"fmt"
	"html/template"
	"interview/domain"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

type CartItemForm struct {
	Product  string `form:"product"   binding:"required"`
	Quantity string `form:"quantity"  binding:"required"`
}

func (pu *cartUsecase) AddItemToCart(c *gin.Context, item string, card string) (domain.CartEntity, error) {
	var isCartNew bool
	var cartEntity domain.CartEntity
	sessionId := item
	_, err := pu.cartRepository.FindBySessionId(c, sessionId)

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.Redirect(302, "/")
			return domain.CartEntity{}, err
		}
		isCartNew = true
		cartEntity = domain.CartEntity{
			SessionID: sessionId,
			Status:    domain.CartOpen,
		}
		pu.cartRepository.Create(c, &cartEntity)
	}

	addItemForm, err := pu.getCartItemForm(c)
	if err != nil {
		c.Redirect(302, "/?error="+err.Error())
		return domain.CartEntity{}, err
	}

	item, ok := itemPriceMapping[addItemForm.Product]
	if !ok {
		c.Redirect(302, "/?error=invalid item name")
		return domain.CartEntity{}, errors.New("Invalid item name")
	}

	quantity, err := strconv.ParseInt(addItemForm.Quantity, 10, 0)
	if err != nil {
		c.Redirect(302, "/?error=invalid quantity")
		return domain.CartEntity{}, err
	}

	var cartItemEntity domain.CartItemEntity
	if isCartNew {
		cartItemEntity = domain.CartItemEntity{
			CartID:      cartdomain.ID,
			ProductName: addItemForm.Product,
			Quantity:    int(quantity),
			Price:       item * float64(quantity),
		}
		pu.cartItemRepository.Create(c, &cartItemEntity)
	} else {
		_, err = pu.cartRepository.FindByProductName(c, cartdomain.ID, addItemForm.Product)

		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				c.Redirect(302, "/")
				return domain.CartEntity{}, err
			}
			cartItemEntity = domain.CartItemEntity{
				CartID:      cartdomain.ID,
				ProductName: addItemForm.Product,
				Quantity:    int(quantity),
				Price:       item * float64(quantity),
			}
			pu.cartItemRepository.Create(c, &cartItemEntity)

		} else {
			cartItemdomain.Quantity += int(quantity)
			cartItemdomain.Price += item * float64(quantity)
			pu.cartItemRepository.Save(c, &cartItemEntity)
		}
	}

	c.Redirect(302, "/")
	return domain.CartEntity{}, nil
}

func (pu *cartUsecase) getCartItemForm(c *gin.Context) (*CartItemForm, error) {
	if c.Request.Body == nil {
		return nil, fmt.Errorf("body cannot be nil")
	}

	form := &CartItemForm{}

	if err := binding.FormPost.Bind(c.Request, form); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return form, nil
}

func (pu *cartUsecase) DeleteCartItem(c *gin.Context, cartId string) error {
	if cartId == "" {
		c.Redirect(302, "/")
		return errors.New("Cart is Empty")
	}

	cookie, _ := c.Request.Cookie("ice_session_id")

	// var cartEntity domain.CartEntity
	_, err := pu.cartRepository.FindBySessionId(c, cookie.Value)
	// result := pu.cartRepository.Where(fmt.Sprintf("status = '%s' AND session_id = '%s'", domain.CartOpen, cookie.Value)).First(&cartEntity)
	if err != nil {
		c.Redirect(302, "/")
		return errors.New("Cart is Incompatible with session value is Empty")
	}

	if cartdomain.Status == domain.CartClosed {
		c.Redirect(302, "/")
		return errors.New("Cart status is CLOSED")
	}

	cartItemID, err := strconv.Atoi(cartItemIDString)
	if err != nil {
		c.Redirect(302, "/")
		return errors.New("Cart is Empty")
	}

	var cartItemEntity domain.CartItemEntity

	_, errCartItem := pu.cartItemRepository.Where(" ID  = ?", cartItemID).First(&cartItemEntity)
	if errCartItem != nil {
		c.Redirect(302, "/")
		return errors.New("Cart is Empty")
	}

	pu.cartItemRepository.Delete(c, &cartItemEntity)
	c.Redirect(302, "/")
	return nil
}

func (pu *cartUsecase) GetCartData(c *gin.Context) domain.CartEntity {
	data := map[string]interface{}{
		"Error": c.Query("error"),
		//"cartItems": cartItems,
	}

	cookie, err := c.Request.Cookie("ice_session_id")
	if err == nil {
		data["CartItems"] = pu.getCartItemData(cookie.Value)
	}

	html, err := pu.renderTemplate(data)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(500)
		return
	}

	c.Header("Content-Type", "text/html")
	c.String(200, html)
}

func (pu *cartUsecase) getCartItemData(sessionID string) (items []map[string]interface{}) {
	var cartEntity domain.CartEntity
	result := pu.cartRepository.Where(fmt.Sprintf("status = '%s' AND session_id = '%s'", domain.CartOpen, sessionID)).First(&cartEntity)

	if result.Error != nil {
		return
	}

	var cartItems []domain.CartItemEntity
	result = pu.cartRepository.Where(fmt.Sprintf("cart_id = %d", cartdomain.ID)).Find(&cartItems)
	if result.Error != nil {
		return
	}

	for _, cartItem := range cartItems {
		item := map[string]interface{}{
			"ID":       cartItem.ID,
			"Quantity": cartItem.Quantity,
			"Price":    cartItem.Price,
			"Product":  cartItem.ProductName,
		}

		items = append(items, item)
	}
	return items
}

func (pu *cartUsecase) renderTemplate(pageData interface{}) (string, error) {
	// Read and parse the HTML template file
	tmpl, err := template.ParseFiles("../../static/add_item_form.html")
	if err != nil {
		return "", fmt.Errorf("Error parsing template: %v ", err)
	}

	// Create a strings.Builder to store the rendered template
	var renderedTemplate strings.Builder

	err = tmpl.Execute(&renderedTemplate, pageData)
	if err != nil {
		return "", fmt.Errorf("Error parsing template: %v ", err)
	}

	// Convert the rendered template to a string
	resultString := renderedTemplate.String()

	return resultString, nil
}
