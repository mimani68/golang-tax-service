package bootstrap

import (
	"interview/config"
	"interview/db"
)

type Application struct {
	Env config.Config
	Db  db.Client
}

func App() Application {
	app := &Application{}
	app.Env = config.GetEnv()
	app.Db = NewMongoDatabase(app.Env.DBConnection)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseDBConnection(app.Db)
}
