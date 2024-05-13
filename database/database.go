package database

import (
	"fmt"
	"k3yst0n3/helpers"
	"k3yst0n3/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbinstanceType struct {
	DB *gorm.DB
}

var DBInstance DbinstanceType

func ConnectDb() {
	dsn := fmt.Sprintf(
		// "host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		"host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC",
		helpers.GetEnv("POSTGRES_USER"),
		helpers.GetEnv("POSTGRES_PASSWORD"),
		helpers.GetEnv("POSTGRES_DB"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Application{})
	db.AutoMigrate(&models.Event{})

	DBInstance = DbinstanceType{
		DB: db,
	}
}
