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

func (pu *cartUsecase) AddItemToCart(c *gin.Context) {
	cookie, _ := c.Request.Cookie("ice_session_id")

	var isCartNew bool
	var cartEntity domain.CartEntity
	result := pu.cartRepository.Where(fmt.Sprintf("status = '%s' AND session_id = '%s'", domain.CartOpen, cookie.Value)) /* .First(&cartEntity) */

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.Redirect(302, "/")
			return
		}
		isCartNew = true
		cartEntity = domain.CartEntity{
			SessionID: cookie.Value,
			Status:    domain.CartOpen,
		}
		pu.cartRepository.Create(c, &cartEntity)
	}

	addItemForm, err := pu.getCartItemForm(c)
	if err != nil {
		c.Redirect(302, "/?error="+err.Error())
		return
	}

	item, ok := itemPriceMapping[addItemForm.Product]
	if !ok {
		c.Redirect(302, "/?error=invalid item name")
		return
	}

	quantity, err := strconv.ParseInt(addItemForm.Quantity, 10, 0)
	if err != nil {
		c.Redirect(302, "/?error=invalid quantity")
		return
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
		result = pu.cartRepository.Where(" cart_id = ? and product_name  = ?", cartdomain.ID, addItemForm.Product).First(&cartItemEntity)

		if result.Error != nil {
			if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.Redirect(302, "/")
				return
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

func (pu *cartUsecase) DeleteCartItem(c *gin.Context) {
	cartItemIDString := c.Query("cart_item_id")
	if cartItemIDString == "" {
		c.Redirect(302, "/")
		return
	}

	cookie, _ := c.Request.Cookie("ice_session_id")

	var cartEntity domain.CartEntity
	result := pu.cartRepository.Where(fmt.Sprintf("status = '%s' AND session_id = '%s'", domain.CartOpen, cookie.Value)).First(&cartEntity)
	if result.Error != nil {
		c.Redirect(302, "/")
		return
	}

	if cartdomain.Status == domain.CartClosed {
		c.Redirect(302, "/")
		return
	}

	cartItemID, err := strconv.Atoi(cartItemIDString)
	if err != nil {
		c.Redirect(302, "/")
		return
	}

	var cartItemEntity domain.CartItemEntity

	result = pu.cartItemRepository.Where(" ID  = ?", cartItemID).First(&cartItemEntity)
	if result.Error != nil {
		c.Redirect(302, "/")
		return
	}

	pu.cartItemRepository.Delete(c, &cartItemEntity)
	c.Redirect(302, "/")
}

func (pu *cartUsecase) GetCartData(c *gin.Context) {
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
