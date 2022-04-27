package routes

import (
	"example/main/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/note")

	//Post note
	api.Post("/", controllers.CreateNote)

	//Get note by id
	api.Get("/", controllers.GetNote)

	//Delete note by id
	api.Delete("/", controllers.DeleteNote)

	// Update note
	api.Put("/", controllers.UpdateNote)

	api.Get("/", controllers.GetAllNotes)

}
