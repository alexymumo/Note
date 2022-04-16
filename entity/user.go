package entity

import "time"

type User struct {
	ID         uint `gorm:"primaryKey"`
	Username   string
	Firstname  string
	Phone      string
	Email      string
	Created_At time.Time
	Updated_At time.Time
}
