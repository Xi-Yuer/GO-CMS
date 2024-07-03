package logic

import (
	"context"
	"fmt"
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
	var user *userModel.User
	if err := l.svcCtx.GormDB.Model(&userModel.User{}).First(&user, "account = ?", in.Account).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &userRPC.GetUserResponse{
		Id:         user.ID,
		Account:    user.Account,
		Password:   "",
		Nickname:   user.NickName,
		Avatar:     user.Avatar,
		Status:     user.Status,
		Department: user.DepartmentID,
		IsAdmin:    user.IsAdmin,
		CreateTime: user.CreatedAt.String(),
		UpdateTime: user.UpdatedAt.String(),
	}, nil
}
