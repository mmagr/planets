package api

import (
	"github.com/mmagr/planets/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Weather struct{
	weatherService service.Weather
}

type output struct {
	Dia int `json:"dia"`
	Clima string`json:"clima"`
}

func NewWeather(weather service.Weather) *Weather {
	return &Weather{weather}
}

func (controller Weather) Register(server *echo.Echo) {
	server.GET("/clima", controller.Render)
}

func (controller Weather) Render(c echo.Context) error {
	dayParam, exists := c.Request().URL.Query()["dia"]
	if !exists {
		return c.JSON(http.StatusBadRequest, map[string]string {
			"mensaje": "parámetro requerido",
			"parametro": "dia",
		})
	}

	// use last supplied value as input
	day, err := strconv.Atoi(dayParam[len(dayParam) - 1])
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string {
			"mensaje": "parámetro invalido: tiene que ser un entero positivo",
			"parametro": "dia",
		})
	}

	if day < 0 {
		return c.JSON(http.StatusBadRequest, map[string]string {
			"mensaje": "parámetro invalido: tiene que ser un entero positivo",
			"parametro": "dia",
		})
	}

	condition, _ := controller.weatherService.ConditionsOn(day)
	result := output{
		Dia: day,
		Clima: condition,
	}

	return c.JSON(http.StatusOK, result)
}
