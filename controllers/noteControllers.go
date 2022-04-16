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
	err := ctx.BodyParser(&note)
	if err != nil {
		return err
	}
	database.DBconn.Create(&note)
	return ctx.JSON(fiber.Map{
		"statusCode": http.StatusOK,
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
	db := database.DBconn
	var notes []entity.Note
	db.Find(&notes)
	return ctx.Status(fiber.StatusOK).JSON(notes)

}

func DeleteNote(ctx *fiber.Ctx) error {
	return nil
}
