package main

import (
	"k3yst0n3/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/health", handlers.HealthCHeckHandler)

	e.Logger.Fatal(e.Start(":4000"))
}
