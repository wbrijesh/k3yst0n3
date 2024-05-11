package main

import (
	"k3yst0n3/handlers"

	"github.com/labstack/echo/v4"
)

func setupRoutes(e *echo.Echo) {
	e.GET("/health-check", handlers.HealthCheckHandler)

	e.GET("/facts", handlers.ListFacts)

	e.POST("/fact", handlers.CreateFact)
}
