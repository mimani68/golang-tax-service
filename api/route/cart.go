package route

import (
	"interview/api/controllers"
	"interview/config"
	"interview/db"
	"time"

	"github.com/gin-gonic/gin"
)

func NewCartRouter(config config.Config, taxController controllers.TaxController, timeout time.Duration, db db.Database, ginEngine *gin.RouterGroup) {
	ginEngine.GET("/", taxController.ShowAddItemForm)
	ginEngine.POST("/add-item", taxController.AddItem)
	ginEngine.GET("/remove-cart-item", taxController.DeleteCartItem)
}
