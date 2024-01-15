package route

import (
	"interview/api/controllers"
	"interview/config"
	"interview/db"
	repository "interview/respoitory"
	"interview/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(config config.Config, timeout time.Duration, db db.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")

	// All Public APIs
	cartRepo := repository.NewCartRepository(db)
	cartItemRepo := repository.NewCartItemRepository(db)
	cartUsecase := usecase.NewCartUsecase(cartRepo, cartItemRepo, timeout)
	taxController := controllers.TaxController{
		Cart: cartUsecase,
	}
	NewCartRouter(config, taxController, timeout, db, publicRouter)

	// All Private APIs
	// NewAdminRouter(config, timeout, db, cartRouter)
}
