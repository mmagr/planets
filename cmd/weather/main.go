package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"

	"github.com/mmagr/planets/internal/api"
	"github.com/mmagr/planets/internal/config"
	"github.com/mmagr/planets/internal/service"
	"github.com/mmagr/planets/internal/util"
)

func main() {
	parsedConfig := config.Init()
	server := echo.New()

	server.Use(middleware.Recover())
	server.Use(middleware.Gzip())
	server.Use(util.MetricsMiddleware())
	server.HideBanner = true
	server.HidePort = true

	p1 := config.Planet(parsedConfig.Sub("planet.vulcano"))
	p2 := config.Planet(parsedConfig.Sub("planet.ferengi"))
	p3 := config.Planet(parsedConfig.Sub("planet.betasoide"))
	api.NewWeather(
		service.NewClimatempo(p1, p2, p3, service.ILineFactory{}, service.TriangleFactory{}),
	).Register(server)

	// Start HTTP server
	go func() {
		log.Info("Starting http server at " + parsedConfig.GetString("server.host"))
		log.Fatal(server.Start(parsedConfig.GetString("server.host")))
	}()

	meta := echo.New()
	meta.HideBanner = true
	meta.HidePort = true
	meta.GET("/prometheus", echo.WrapHandler(promhttp.Handler()))
	api.NewHealth().Register(meta)

	go func() {
		log.Info("Starting metadata server at " + parsedConfig.GetString("meta.host"))
		log.Fatal(meta.Start(parsedConfig.GetString("meta.host")))
	}()

	// Listen for system signals to gracefully stop the application
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	switch <-signalChannel {
	case os.Interrupt:
		log.Info("Received SIGINT, stopping...")
	case syscall.SIGTERM:
		log.Info("Received SIGTERM, stopping...")
	}
}
