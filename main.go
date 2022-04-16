package main

import (
	"example/main/database"
	"example/main/entity"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Database connected successfully")
	database.DBconn.AutoMigrate(&entity.Note{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	app.Listen(":3000")

}
