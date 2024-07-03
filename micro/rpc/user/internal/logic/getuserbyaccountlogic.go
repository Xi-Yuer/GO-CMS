package logic

import (
	"context"
	userModel "micro/model/user"

	"micro/rpc/user/internal/svc"
	"micro/rpc/user/userRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByAccountLogic {
	return &GetUserByAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByAccountLogic) GetUserByAccount(in *userRPC.GetUserByAccountRequest) (*userRPC.GetUserResponse, error) {
	var user *userRPC.GetUserResponse
	if err := l.svcCtx.GormDB.Model(&userModel.User{}).First(&userModel.User{}, "account = ?", in.Account).Error; err != nil {
		return nil, err
	}
	return user, nil
}
