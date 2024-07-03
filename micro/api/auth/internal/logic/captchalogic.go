package logic

import (
	"context"
	"micro/rpc/captcha/captcha"
	"micro/shared/snowflake"
	"strconv"

	"micro/api/auth/internal/svc"
	"micro/api/auth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaLogic) Captcha(req *types.EmptyRequest) (resp *types.CommonResponse, err error) {
	sessionID := strconv.FormatInt(snowflake.GenID(), 10)
	generateCaptcha, err := l.svcCtx.CaptchaService.GenerateCaptcha(l.ctx, &captcha.GenerateCaptchaRequest{
		SessionId: sessionID,
	})

	if err != nil {
		return nil, err
	}

	return &types.CommonResponse{
		Code: 0,
		Data: &struct {
			Captcha     string `json:"captchaCode"`
			SessionID   string `json:"sessionID"`
			CaptchaCode string `json:"captchaBase64Code"`
		}{
			CaptchaCode: generateCaptcha.CaptchaCode,
			Captcha:     generateCaptcha.Captcha,
			SessionID:   sessionID,
		},
		Msg: "success",
	}, nil
}
