package routes

import (
	"example/main/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("api")
	api.Post("save", controllers.CreateNote)

}
