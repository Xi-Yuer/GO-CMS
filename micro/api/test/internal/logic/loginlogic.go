package logic

import (
	"context"
	"micro/shared/token"

	"micro/api/test/internal/svc"
	"micro/api/test/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	createToken, err := token.CreateToken(100, "xiyuer")
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		ID:   "1",
		Name: createToken,
	}, nil
}
