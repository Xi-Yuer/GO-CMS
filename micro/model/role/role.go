package roleModlel

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID          string `gorm:"<-:create;primaryKey"`
	RoleName    string `gorm:"not null"`
	Description string
	CanEdit     bool `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
