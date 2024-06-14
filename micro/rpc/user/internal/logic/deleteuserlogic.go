package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &userRPC.CommonResponse{}, nil
}
