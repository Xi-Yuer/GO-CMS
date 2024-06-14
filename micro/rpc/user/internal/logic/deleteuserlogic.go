package logic

import (
	"context"
	userModel "micro/model/user"

	"micro/rpc/user/internal/svc"
	"micro/rpc/user/userRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *userRPC.DeleteUserRequest) (*userRPC.CommonResponse, error) {
	if err := l.svcCtx.GormDB.Model(&userModel.User{}).Delete(in.Id).Error; err != nil {
		return nil, err
	}
	return &userRPC.CommonResponse{
		Ok:  true,
		Msg: "删除成功",
	}, nil
}
