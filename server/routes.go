package main

import (
	"k3yst0n3/internal/auth"

	"github.com/labstack/echo/v4"
)

func setupRoutes(e *echo.Echo) {
	e.GET("/health-check", HealthCheckHandler)

	// authentication routes
	e.POST("/register", auth.RegisterUser)
	e.POST("/login", auth.LoginUser)
}
