package main

import (
	"github.com/ammorteza/location_history_server/app"
	"github.com/ammorteza/location_history_server/http"
)

func main() {
	app := app.New()

	server := http.NewServer(app)
	server.Start()
}
