package handlers

import (
	"time"

	"yourmodule/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func createToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func LoginHandler(c *fiber.Ctx) error {
	var u models.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	if u.Username == "Chek" && u.Password == "123456" {
		tokenString, err := createToken(u.Username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Could not create token")
		}
		return c.Status(fiber.StatusOK).SendString(tokenString)
	}
	return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
}

func ProtectedHandler(c *fiber.Ctx) error {
	return c.SendString("Welcome to the protected area")
}
