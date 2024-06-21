package userModel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	Account      string `gorm:"unique"`
	Password     string `gorm:"not null"`
	NickName     string
	Avatar       string
	Status       bool `gorm:"default:true"`
	DepartmentID string
	IsAdmin      bool `gorm:"default:true"`
}
