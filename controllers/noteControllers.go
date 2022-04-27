package controllers

import (
	"example/main/database"
	"example/main/entity"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

// save note
func CreateNote(ctx *fiber.Ctx) error {
	db := database.DBconn
	note := new(entity.Note)
	err := ctx.BodyParser(note)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	note.ID = uuid.New()
	err = db.Create(&note).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create note", "data": err})
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": note})

}

// Get note by id
func GetNote(c *fiber.Ctx) error {
	db := database.DBconn
	var note entity.Note
	id := c.Params("id")
	db.Find(&note, "id = ?", id)
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Note found", "data": note})
}

// Get all notes
func GetAllNotes(c *fiber.Ctx) error {
	return nil
}

// Delete note by id
func DeleteNote(c *fiber.Ctx) error {
	return nil
}

// Update note
func UpdateNote(c *fiber.Ctx) error {
	return nil
}
