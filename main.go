package main

import (
	"log"

	"go-fiber-jwt-auth/handlers"
	"go-fiber-jwt-auth/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Go Fiber JWT Authentication API")
	})

	app.Post("/login", handlers.LoginHandler)
	app.Get("/protected", middleware.Protected(), handlers.ProtectedHandler)

	log.Fatal(app.Listen(":4000"))
}
