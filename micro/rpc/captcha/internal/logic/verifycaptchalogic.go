package logic

import (
	"context"
	"errors"
	"fmt"

	"micro/rpc/captcha/captcha"
	"micro/rpc/captcha/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCaptchaLogic {
	return &VerifyCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyCaptchaLogic) VerifyCaptcha(in *captcha.VerifyCaptchaRequest) (*captcha.VerifyCaptchaResponse, error) {
	sessionID := in.GetSessionId()
	userInputCode := in.GetCaptchaCode()

	// 从 Redis 中获取存储的验证码
	storedCode, err := l.svcCtx.Redis.Get(l.ctx, fmt.Sprintf("captcha:%s", sessionID)).Result()
	if err != nil {
		return nil, errors.New("验证码已过期，请重新获取")
	}

	// 检查存储的验证码是否与用户输入的验证码匹配
	valid := storedCode == userInputCode

	// 删除 Redis 中存储的验证码
	l.svcCtx.Redis.Del(l.ctx, fmt.Sprintf("captcha:%s", sessionID))

	return &captcha.VerifyCaptchaResponse{
		Valid: valid,
	}, nil
}
