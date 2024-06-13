package logic

import (
	"context"
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
		return nil, err
	}

	// 检查存储的验证码是否与用户输入的验证码匹配
	valid := storedCode == userInputCode

	return &captcha.VerifyCaptchaResponse{
		Valid: valid,
	}, nil
}
