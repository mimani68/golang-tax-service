package main

import (
	"interview/api/route"
	"interview/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDBConnection()

	gin := gin.Default()
	timeout := time.Duration(10000) * time.Second
	route.Setup(app.Env, timeout, app.Db, gin)
	gin.Run(app.Env.ServerAddress)
}
