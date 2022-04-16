package entity

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Created_At  time.Time
	Updated_At  time.Time
	Deleted_At  time.Time
}
