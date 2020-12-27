package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Health struct{}

func NewHealth() *Health {
	return &Health{}
}

func (controller Health) Register(server *echo.Echo) {
	server.GET("/health", controller.Fetch)
}

func (controller Health) Fetch(c echo.Context) error {

	response := make(map[string]string)
	response["status"] = "UP"
	return c.JSON(http.StatusOK, response)
}
