package http

import (
	"fmt"
	"log"

	"github.com/ammorteza/location_history_server/app"
	"github.com/ammorteza/location_history_server/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	httpServer *echo.Echo
}

func NewServer(app *app.Application) *Server {
	// initiate new http server
	server := &Server{}
	server.httpServer = echo.New()
	server.httpServer.Use(middleware.Logger())
	server.httpServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// define new routes
	server.initRoute(app)

	return server
}

func (s *Server) Start() {
	log.Fatal(s.httpServer.Start(fmt.Sprintf(":%s", config.Conf.Server.Port)))
}
