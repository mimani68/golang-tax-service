package route

import (
	"interview/api/controllers"
	"interview/config"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewCartRouter(config config.Config, taxController controllers.CartController, timeout time.Duration, db *gorm.DB, ginEngine *gin.RouterGroup) {
	ginEngine.GET("/", taxController.ShowAddItemForm)
	ginEngine.POST("/add-item", taxController.AddItem)
	ginEngine.GET("/remove-cart-item", taxController.DeleteCartItem)
}
