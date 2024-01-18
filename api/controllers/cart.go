package controllers

import (
	"errors"
	"interview/domain"
	"interview/pkg/html"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type CartController struct {
	Cart domain.CartUsecase
}

func (t *CartController) ShowAddItemForm(c *gin.Context) {
	data := map[string]interface{}{
		"Error": c.Query("error"),
	}

	_, err := c.Request.Cookie("ice_session_id")
	if errors.Is(err, http.ErrNoCookie) {
		c.SetCookie("ice_session_id", time.Now().String(), 3600, "/", "localhost", false, true)
	}

	cookie, err := c.Request.Cookie("ice_session_id")
	if err == nil {
		data["CartItems"], _ = t.Cart.GetCartItemData(c, cookie.Value)
	}

	html, err := html.RenderTemplate("../static/add_item_form.html", data)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(400)
		return
	}

	c.Header("Content-Type", "text/html")
	c.String(200, html)
}

func (t *CartController) AddItem(c *gin.Context) {
	cookie, err := c.Request.Cookie("ice_session_id")

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		c.Redirect(302, "/")
		return
	}

	if c.Request.Body == nil {
		c.Redirect(400, "/")
		return
	}

	itemObject := &domain.CartItemForm{}

	if err := binding.FormPost.Bind(c.Request, itemObject); err != nil {
		log.Println(err.Error())
		c.Redirect(302, "?error="+err.Error())
		return
	}

	err = t.Cart.AddItemToCart(c, cookie.Value, *itemObject)
	if err != nil {
		c.Redirect(302, "/?error="+err.Error())
	} else {
		c.Redirect(302, "/")
	}
}

func (t *CartController) DeleteCartItem(c *gin.Context) {
	cartItemIDString := c.Query("cart_item_id")
	if cartItemIDString == "" {
		c.Redirect(302, "/")
		return
	}

	cookie, err := c.Request.Cookie("ice_session_id")

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		c.Redirect(302, "/")
		return
	}

	err = t.Cart.DeleteCartItem(c, cartItemIDString)
	if err != nil {
		c.Redirect(302, "/")
	} else {
		c.Redirect(302, "/?error="+err.Error())
	}
}
