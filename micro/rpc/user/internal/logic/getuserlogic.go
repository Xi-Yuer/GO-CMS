package logic

import (
	"context"
	userModel "micro/model/user"

	"micro/rpc/user/internal/svc"
	"micro/rpc/user/userRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *userRPC.GetUserRequest) (*userRPC.GetUserResponse, error) {
	var user userModel.User
	if err := l.svcCtx.GormDB.Model(&userModel.User{}).First(&user, in.Id).Error; err != nil {
		return nil, err
	}
	return &userRPC.GetUserResponse{
		Id:         user.ID,
		Account:    user.Account,
		Nickname:   user.NickName,
		Avatar:     user.Avatar,
		Department: user.DepartmentID,
		Status:     user.Status,
		IsAdmin:    user.IsAdmin,
		CreateTime: user.CreatedAt.String(),
		UpdateTime: user.UpdatedAt.String(),
	}, nil
}
