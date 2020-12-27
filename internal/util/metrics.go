package util

import (
	echo "github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"reflect"
	"time"
)

func normalizeHTTPStatus(status int) string {
	if status < 200 {
		return "1xx"
	} else if status < 300 {
		return "2xx"
	} else if status < 400 {
		return "3xx"
	} else if status < 500 {
		return "4xx"
	}
	return "5xx"
}

func isNotFoundHandler(handler echo.HandlerFunc) bool {
	return reflect.ValueOf(handler).Pointer() == reflect.ValueOf(echo.NotFoundHandler).Pointer()
}

func MetricsMiddleware() echo.MiddlewareFunc {

	httpSummary := promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "http_request",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"status", "method", "path"})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			path := c.Path()

			// to avoid attack high cardinality of 404
			if isNotFoundHandler(c.Handler()) {
				path = "-unknown-"
			}

			start := time.Now()
			err := next(c)
			end := time.Now()

			if err != nil {
				c.Error(err)
			}

			httpSummary.WithLabelValues(
				normalizeHTTPStatus(c.Response().Status),
				req.Method, path).Observe(end.Sub(start).Seconds())
			return err
		}
	}
}
