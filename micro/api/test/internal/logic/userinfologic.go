package logic

import (
	"context"
	"fmt"

	"micro/api/test/internal/svc"
	"micro/api/test/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	value := l.ctx.Value("user_id")
	fmt.Println(value)
	return &types.UserInfoResp{
		Name: "哈哈",
	}, nil
}
