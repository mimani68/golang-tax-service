package route

import (
	"interview/config"
	"interview/db"
	"time"

	"github.com/gin-gonic/gin"
)

func NewCartRouter(config config.Config, timeout time.Duration, db db.Database, group *gin.RouterGroup) {
	// ur := repository.NewUserRepository(db, domain.CollectionUser)
	// pc := &controller.ProfileController{
	// 	ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	// }
	// group.GET("/profile", pc.Fetch)
}
