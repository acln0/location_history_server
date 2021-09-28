package http

import (
	"github.com/ammorteza/location_history_server/app"
	"github.com/ammorteza/location_history_server/http/handler"
)

func (s *Server) initRoute(app *app.Application) {
	location := s.httpServer.Group("/location")
	{
		location.POST("/:order_id/now", handler.NewLocation(app))
		location.GET("/:order_id/:max", handler.GetHistoryOfLocation(app))
		location.DELETE("/:order_id", handler.DeleteHistory(app))
	}
}
