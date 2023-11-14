package main

import (
	"log"

	"github.com/artisbasecode/api-fiber-gorm/database"
	"github.com/artisbasecode/api-fiber-gorm/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to this awesome api")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)

	// User

	app.Post("/api/users", routes.CreateUser)
}

func main() {

	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
