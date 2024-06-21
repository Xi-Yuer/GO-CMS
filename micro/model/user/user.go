package userModel

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           string `gorm:"<-:create;primaryKey"`
	Account      string `gorm:"<-:create;unique"`
	Password     string `gorm:"not null"`
	NickName     string
	Avatar       string
	Status       bool `gorm:"default:true"`
	DepartmentID string
	IsAdmin      bool `gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
