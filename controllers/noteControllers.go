package controllers

import (
	"example/main/database"
	"example/main/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*
* function to create a note
 */

func CreateNote(ctx *fiber.Ctx) error {
	var note entity.Note
	if err := ctx.BodyParser(&note); err != nil {
		return err
	}
	database.DBconn.Create(&note)
	return ctx.JSON(fiber.Map{
		"StatusCode": http.StatusOK,
		"message":    "Created Successfully",
	})
}

func GetNoteByID(ctx *fiber.Ctx) error {
	return nil
}

/*
* @func GetNotes
* Get all notes
* fiber context
 */

func GetNotes(ctx *fiber.Ctx) error {
	return nil
}

func DeleteNote(ctx *fiber.Ctx) error {
	return nil
}
