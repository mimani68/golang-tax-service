package route

import (
	"interview/api/controllers"
	"interview/config"
	"interview/db"
	"time"

	"github.com/gin-gonic/gin"
)

func NewCartRouter(config config.Config, timeout time.Duration, db db.Database, ginEngine *gin.RouterGroup) {
	// ur := repository.NewUserRepository(db, domain.CollectionUser)
	// pc := &controller.ProfileController{
	// 	ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	// }
	// group.GET("/profile", pc.Fetch)
	var taxController controllers.TaxController
	ginEngine.GET("/", taxController.ShowAddItemForm)
	ginEngine.POST("/add-item", taxController.AddItem)
	ginEngine.GET("/remove-cart-item", taxController.DeleteCartItem)

}
