package models

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type JwtCustomClaims struct {
	jwt.RegisteredClaims
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	ID        uint   `json:"id"`
}

type User struct {
	gorm.Model
	FirstName    string        `json:"first_name" gorm:"text"`
	LastName     string        `json:"last_name" gorm:"text"`
	Email        string        `json:"email" gorm:"text"`
	Password     string        `json:"password" gorm:"text"`
	Role         string        `json:"role" gorm:"text"`
	Applications []Application `json:"applications" gorm:"foreignKey:UserID"`
}

type Application struct {
	gorm.Model
	UserID       string  `json:"user_id" gorm:"text"`
	Name         string  `json:"name" gorm:"text"`
	Description  string  `json:"description" gorm:"text"`
	Domain       string  `json:"domain" gorm:"text"`
	Status       string  `json:"status" gorm:"text"`
	TrackingCode string  `json:"tracking_code" gorm:"text"`
	Events       []Event `json:"events" gorm:"foreignKey:ApplicationID"`
}

type Event struct {
	gorm.Model
	ApplicationID string `json:"application_id" gorm:"text"`
	Name          string `json:"name" gorm:"text"`
	Description   string `json:"description" gorm:"text"`
}
