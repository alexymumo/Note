package entity

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID          int    `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
