package app

import "github.com/ammorteza/location_history_server/db"

type Application struct {
	Storage *db.Storage
}

func New() *Application {
	return &Application{
		Storage: db.New(),
	}
}
