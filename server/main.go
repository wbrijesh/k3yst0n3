package main

import (
	"k3yst0n3/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.ConnectDb()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())

	setupRoutes(e)

	e.Logger.Fatal(e.Start(":4000"))
}
