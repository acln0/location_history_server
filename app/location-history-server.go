package app

import (
	"github.com/ammorteza/location_history_server/config"
	"github.com/ammorteza/location_history_server/db"
)

type Application struct {
	Storage *db.Storage
	Config  config.Config
}

func New() *Application {
	config.InitConfig(".env")
	return &Application{
		Config:  config.Conf,
		Storage: db.New(),
	}
}
