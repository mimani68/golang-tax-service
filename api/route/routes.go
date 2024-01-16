package route

import (
	"interview/api/controllers"
	"interview/config"
	repository "interview/respoitory"
	"interview/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(config config.Config, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")

	// All Public APIs
	cartRepo := repository.NewCartRepository(db)
	cartItemRepo := repository.NewCartItemRepository(db)
	cartUsecase := usecase.NewCartUsecase(cartRepo, cartItemRepo, timeout)
	taxController := controllers.CartController{
		Cart: cartUsecase,
	}
	NewCartRouter(config, taxController, timeout, db, publicRouter)

	// All Private APIs
	// NewAdminRouter(config, timeout, db, cartRouter)
}
