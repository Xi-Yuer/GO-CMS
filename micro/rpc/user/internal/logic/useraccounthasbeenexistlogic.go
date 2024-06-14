package logic

import (
	"context"
	userModel "micro/model/user"

	"micro/rpc/user/internal/svc"
	"micro/rpc/user/userRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAccountHasBeenExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAccountHasBeenExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAccountHasBeenExistLogic {
	return &UserAccountHasBeenExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAccountHasBeenExistLogic) UserAccountHasBeenExist(in *userRPC.UserAccountHasBeenExistRequest) (*userRPC.CommonResponse, error) {
	find := l.svcCtx.GormDB.Model(&userModel.User{}).Find(&userModel.User{}, "account = ?", in.Account)
	if find.RowsAffected > 0 {
		return &userRPC.CommonResponse{
			Ok:  true,
			Msg: "账号已存在",
		}, nil
	}

	return &userRPC.CommonResponse{
		Ok:  false,
		Msg: "账号不存在",
	}, nil
}
