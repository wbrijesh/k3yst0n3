package main

import (
	"k3yst0n3/database"

	"github.com/labstack/echo/v4"
)

func main() {
	database.ConnectDb()

	e := echo.New()

	setupRoutes(e)

	e.Logger.Fatal(e.Start(":4000"))
}
