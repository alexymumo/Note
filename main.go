package main

import (
	"example/main/controllers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	api := app.Group("/api")
	SetUpRoutes(api)
}

func SetUpRoutes(group fiber.Router) {
	noteRoute := group.Group("/notes")
	noteRoute.Post("/", controllers.CreateNote)
}

func main() {
	app := fiber.New()
	SetUp(app)
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running")

}
