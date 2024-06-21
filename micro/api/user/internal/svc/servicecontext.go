package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"micro/api/user/internal/config"
	"micro/rpc/user/userservice"
)

type ServiceContext struct {
	Config      config.Config
	UserService userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserService: userservice.NewUserService(zrpc.MustNewClient(c.UserService)),
	}
}
