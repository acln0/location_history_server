package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ammorteza/location_history_server/app"
	"github.com/ammorteza/location_history_server/entity"
	"github.com/labstack/echo/v4"
)

type locationResponse struct {
	OrderID string            `json:"order_id"`
	History []entity.Location `json:"history"`
}

func NewLocation(app *app.Application) echo.HandlerFunc {
	return func(c echo.Context) error {
		var location entity.Location

		if err := c.Bind(&location); err != nil {
			return err
		}

		orderID := c.Param("order_id")
		if orderID == "" || location.Lat == "" || location.Lng == "" {
			return c.NoContent(http.StatusBadRequest)
		}

		app.Storage.Insert(orderID, location)

		return c.NoContent(http.StatusOK)
	}
}

func GetHistoryOfLocation(app *app.Application) echo.HandlerFunc {
	return func(c echo.Context) error {
		var result locationResponse
		orderID := c.Param("order_id")
		if orderID == "" {
			return c.NoContent(http.StatusBadRequest)
		}

		max, err := strconv.Atoi(c.Param("max"))
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		log.Println(orderID)
		log.Println(max)

		result.OrderID = orderID
		result.History = app.Storage.Fetch(orderID, max)

		return c.JSON(http.StatusOK, result)
	}
}

func DeleteHistory(app *app.Application) echo.HandlerFunc {
	return func(c echo.Context) error {
		orderID := c.Param("order_id")
		if orderID == "" {
			return c.NoContent(http.StatusBadRequest)
		}

		app.Storage.Delete(orderID)
		return c.NoContent(http.StatusOK)
	}
}
