package logic

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"micro/rpc/captcha/captcha"
	"micro/rpc/captcha/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateCaptchaLogic {
	return &GenerateCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateCaptchaLogic) GenerateCaptcha(in *captcha.GenerateCaptchaRequest) (*captcha.GenerateCaptchaResponse, error) {
	sessionID := in.GetSessionId()
	// 生成一个随机的 6 位验证码
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	// 设置验证码的过期时间为 5 分钟后
	expiresAt := time.Now().Add(5 * time.Minute).Unix()

	// 将验证码存储到 Redis 中，以 sessionID 为键
	err := l.svcCtx.Redis.Set(l.ctx, fmt.Sprintf("captcha:%s", sessionID), code, time.Duration(expiresAt)*time.Second).Err()
	if err != nil {
		return nil, err
	}

	return &captcha.GenerateCaptchaResponse{
		CaptchaCode: code,
		ExpiresAt:   expiresAt,
	}, nil
}
