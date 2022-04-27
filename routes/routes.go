package routes

import (
	"example/main/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/note")

	//Post note
	api.Post("/", controllers.CreateNote)

	//Get note
	api.Get("/", controllers.GetNote)

}
