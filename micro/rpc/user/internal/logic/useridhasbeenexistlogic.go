package logic

import (
	"context"
	userModel "micro/model/user"

	"micro/rpc/user/internal/svc"
	"micro/rpc/user/userRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserIDHasBeenExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserIDHasBeenExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserIDHasBeenExistLogic {
	return &UserIDHasBeenExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserIDHasBeenExistLogic) UserIDHasBeenExist(in *userRPC.DeleteUserRequest) (*userRPC.CommonResponse, error) {
	find := l.svcCtx.GormDB.Model(&userModel.User{}).Find(&userModel.User{}, "id = ?", in.Id)
	if find.RowsAffected > 0 {
		return &userRPC.CommonResponse{
			Ok:  true,
			Msg: "",
		}, nil
	}

	return &userRPC.CommonResponse{
		Ok:  false,
		Msg: "用户不存在",
	}, nil
}
