package route

import (
	"interview/config"
	"interview/db"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(config config.Config, timeout time.Duration, db db.Database, gin *gin.Engine) {
	cartRouter := gin.Group("")

	// All Public APIs
	NewCartRouter(config, timeout, db, cartRouter)

	// All Private APIs
	// NewAdminRouter(config, timeout, db, cartRouter)
}
