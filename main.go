package main

import (
	"example/main/database"
	"example/main/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":3000")
}
