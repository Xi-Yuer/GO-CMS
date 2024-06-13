package logic

import (
	"context"
	"micro/rpc/captcha/captcha"

	"micro/api/auth/internal/svc"
	"micro/api/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.CommonResponse, err error) {
	verifyCaptcha, err := l.svcCtx.CaptchaService.VerifyCaptcha(l.ctx, &captcha.VerifyCaptchaRequest{
		SessionId: req.SessionID,
	})
	if err != nil {
		return nil, err
	}

	return &types.CommonResponse{
		Code: 0,
		Data: verifyCaptcha,
		Msg:  "",
	}, nil
}
