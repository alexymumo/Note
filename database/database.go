package database

import (
	"example/main/entity"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBconn *gorm.DB

func initDatabase() {
	var err error

	// connect to sqlite
	DBconn, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("Database connected successfully")

}

func AutoMigrate() {
	DBconn.AutoMigrate(&entity.Note{})
	fmt.Println("Database migrated")
}
