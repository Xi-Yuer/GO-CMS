package logic

import (
	"context"
	"micro/api/auth/internal/svc"
	"micro/api/auth/internal/types"
	"micro/rpc/captcha/captcha"
	"micro/rpc/user/userRPC"
	"micro/shared/token"

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
		SessionId:   req.SessionID,
		CaptchaCode: req.Captcha,
	})
	if err != nil {
		return &types.CommonResponse{
			Code: 400,
			Data: nil,
			Msg:  err.Error(),
		}, nil
	}
	if !verifyCaptcha.Valid {
		return &types.CommonResponse{
			Code: 400,
			Data: nil,
			Msg:  "验证码错误",
		}, nil
	}

	hasBeenExist, err := l.svcCtx.UserService.UserAccountHasBeenExist(l.ctx, &userRPC.UserAccountHasBeenExistRequest{Account: req.Account})
	if err != nil {
		return nil, err
	}

	if !hasBeenExist.Ok {
		return &types.CommonResponse{
			Code: 400,
			Data: nil,
			Msg:  "用户不存在",
		}, nil
	}

	userResponse, err := l.svcCtx.UserService.GetUserByAccount(l.ctx, &userRPC.GetUserByAccountRequest{Account: req.Account})
	if err != nil {
		return &types.CommonResponse{
			Code: 400,
			Data: nil,
			Msg:  "用户不存在",
		}, nil
	}

	createToken, err := token.CreateToken(userResponse.Id, userResponse.Account)

	if err != nil {
		return &types.CommonResponse{
			Code: 400,
			Data: nil,
			Msg:  err.Error(),
		}, nil
	}

	return &types.CommonResponse{
		Code: 200,
		Data: createToken,
		Msg:  "success",
	}, nil
}
