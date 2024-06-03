package main

import (
	"log"

	"github.com/aselasperera/go-fiber-jwt-auth/handlers"
	"github.com/aselasperera/go-fiber-jwt-auth/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/login", handlers.LoginHandler)
	app.Get("/protected", middleware.Protected(), handlers.ProtectedHandler)

	log.Fatal(app.Listen(":4000"))
}
