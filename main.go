package main

import (
	"example/main/database"
	"example/main/routes"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	routes.Setup(app)
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
