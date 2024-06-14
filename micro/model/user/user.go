package userModel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           string
	Account      string
	Password     string
	NickName     string
	Avatar       string
	Status       bool
	DepartmentID string
	IsAdmin      bool
}
