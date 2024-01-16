package controllers

import (
	"errors"
	"interview/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	Cart domain.CartUsecase
}

func (t *CartController) ShowAddItemForm(c *gin.Context) {
	_, err := c.Request.Cookie("ice_session_id")
	if errors.Is(err, http.ErrNoCookie) {
		c.SetCookie("ice_session_id", time.Now().String(), 3600, "/", "localhost", false, true)
	}

	t.Cart.GetCartData(c)
}

func (t *CartController) AddItem(c *gin.Context) {
	cookie, err := c.Request.Cookie("ice_session_id")

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		c.Redirect(302, "/")
		return
	}

	t.Cart.AddItemToCart(c, "", "")
}

func (t *CartController) DeleteCartItem(c *gin.Context) {
	cookie, err := c.Request.Cookie("ice_session_id")

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		c.Redirect(302, "/")
		return
	}

	t.Cart.DeleteCartItem(c, "")
}
