package controllers

import (
	"example/main/database"

	"github.com/gofiber/fiber/v2"
)

func CreateNote(ctx *fiber.Ctx) error {
	db := database.DBconn
	type request struct {
	}
	var body request 
	return nil
}

func GetNoteByID(ctx *fiber.Ctx) error {
	return nil

}

func GetNotes(ctx *fiber.Ctx) error {
	return nil
}

func DeleteNote(ctx *fiber.Ctx) error {
	return nil
}
