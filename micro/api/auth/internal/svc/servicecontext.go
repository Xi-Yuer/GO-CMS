package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"micro/api/auth/internal/config"
	"micro/rpc/captcha/captchaservice"
)

type ServiceContext struct {
	Config         config.Config
	CaptchaService captchaservice.CaptchaService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CaptchaService: captchaservice.NewCaptchaService(zrpc.MustNewClient(c.CaptchaService)),
	}
}
