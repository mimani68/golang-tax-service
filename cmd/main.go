package main

import (
	"interview/api/route"
	"interview/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	db := app.Db.Database("sample")
	defer app.CloseDBConnection()

	gin := gin.Default()
	timeout := time.Duration(10000) * time.Second
	route.Setup(app.Env, timeout, db, gin)
	gin.Run(app.Env.ServerAddress)

}
