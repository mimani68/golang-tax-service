package bootstrap

import (
	"interview/config"

	"gorm.io/gorm"
)

type Application struct {
	Env config.Config
	Db  *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = config.GetEnv()
	app.Db = NewDatabase(app.Env.DBConnection)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseDBConnection(app.Db)
}
