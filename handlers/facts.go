package handlers

import (
	"k3yst0n3/database"
	"k3yst0n3/models"

	"github.com/labstack/echo/v4"
)

func ListFacts(c echo.Context) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)
	return c.JSON(200, facts)
}

func CreateFact(c echo.Context) error {
	fact := new(models.Fact)
	if err := c.Bind(fact); err != nil {
		return c.JSON(500, map[string]string{"message": err.Error()})
	}
	database.DB.Db.Create(&fact)
	return c.JSON(200, fact)
}
