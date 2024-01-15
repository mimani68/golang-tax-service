package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"interview/usecase/calculator"
	"net/http"
	"time"
)

type TaxController struct {
	calculator: calculator
}

func (t *TaxController) ShowAddItemForm(c *gin.Context) {
	_, err := c.Request.Cookie("ice_session_id")
	if errors.Is(err, http.ErrNoCookie) {
		c.SetCookie("ice_session_id", time.Now().String(), 3600, "/", "localhost", false, true)
	}

	t.calculator.GetCartData(c)
}

func (t *TaxController) AddItem(c *gin.Context) {
	cookie, err := c.Request.Cookie("ice_session_id")

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		c.Redirect(302, "/")
		return
	}

	t.calculator.AddItemToCart(c)
}

func (t *TaxController) DeleteCartItem(c *gin.Context) {
	cookie, err := c.Request.Cookie("ice_session_id")

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		c.Redirect(302, "/")
		return
	}

	t.calculator.DeleteCartItem(c)
}
