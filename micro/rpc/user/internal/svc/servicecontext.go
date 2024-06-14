package svc

import (
	"gorm.io/gorm"
	"micro/rpc/user/internal/config"
	gormDB "micro/shared/gorm"
)

type ServiceContext struct {
	Config config.Config
	GormDB *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		GormDB: gormDB.NewGorm(),
	}
}
