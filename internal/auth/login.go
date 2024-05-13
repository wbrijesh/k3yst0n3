package auth

import (
	"k3yst0n3/database"
	"k3yst0n3/helpers"
	"k3yst0n3/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func LoginUser(c echo.Context) error {
	requestID := c.Response().Header().Get(echo.HeaderXRequestID)
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(500, map[string]string{
			"error":      "error binding user from request: " + err.Error(),
			"request_id": requestID,
		})
	}
	necessaryFields := map[string]interface{}{
		"email":    user.Email,
		"password": user.Password,
	}
	for key, field := range necessaryFields {
		if field == "" {
			return c.JSON(400, map[string]string{
				"error":      key + " is required",
				"request_id": requestID,
			})
		}
	}
	var existingUser models.User
	database.DBInstance.DB.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.Email != user.Email {
		return c.JSON(400, map[string]string{
			"error":      "User does not exist",
			"request_id": requestID,
		})
	}
	if !helpers.CheckPasswordHash(user.Password, existingUser.Password) {
		return c.JSON(400, map[string]string{
			"error":      "Invalid password",
			"request_id": requestID,
		})
	}
	claims := &JwtCustomClaims{
		ID:        existingUser.ID,
		FirstName: existingUser.FirstName,
		LastName:  existingUser.LastName,
		Role:      existingUser.Role,
		Email:     existingUser.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "k3yst0n3",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	encodedToken, err := token.SignedString([]byte(helpers.GetEnv("JWT_SIGNING_KEY")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":      "Error while generating token",
			"request_id": requestID,
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message":    "User logged in successfully",
		"token":      encodedToken,
		"request_id": requestID,
	})
}
