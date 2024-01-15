package route

import (
	"interview/api/controllers"
	"interview/config"
	"interview/db"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(config config.Config, timeout time.Duration, db db.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")

	// All Public APIs
	var taxController controllers.TaxController
	NewCartRouter(config, taxController, timeout, db, publicRouter)

	// All Private APIs
	// NewAdminRouter(config, timeout, db, cartRouter)
}
