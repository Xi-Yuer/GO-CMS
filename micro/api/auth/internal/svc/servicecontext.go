package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"micro/api/auth/internal/config"
	"micro/rpc/captcha/captchaservice"
	"micro/rpc/user/userservice"
)

type ServiceContext struct {
	Config         config.Config
	CaptchaService captchaservice.CaptchaService
	UserService    userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CaptchaService: captchaservice.NewCaptchaService(zrpc.MustNewClient(c.CaptchaService)),
		UserService:    userservice.NewUserService(zrpc.MustNewClient(c.UserService)),
	}
}
