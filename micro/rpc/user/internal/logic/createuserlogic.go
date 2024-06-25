package logic

import (
	"context"
	userModel "micro/model/user"

	"micro/rpc/user/internal/svc"
	"micro/rpc/user/userRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *userRPC.CreateUserRequest) (*userRPC.CommonResponse, error) {
	beenExistLogic := NewUserAccountHasBeenExistLogic(l.ctx, l.svcCtx)
	exist, err := beenExistLogic.UserAccountHasBeenExist(&userRPC.UserAccountHasBeenExistRequest{
		Account: in.Account,
	})
	if err != nil {
		return &userRPC.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		}, err
	}
	if exist.Ok {
		return &userRPC.CommonResponse{
			Ok:  false,
			Msg: "用户名已存在",
		}, nil
	}
	if err := l.svcCtx.GormDB.Create(&userModel.User{
		ID:           in.Id,
		Account:      in.Account,
		Password:     in.Password,
		Avatar:       in.Avatar,
		NickName:     in.Nickname,
		DepartmentID: in.Department,
		Status:       in.Status,
		IsAdmin:      in.IsAdmin,
	}).Error; err != nil {
		return &userRPC.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		}, nil
	}
	return &userRPC.CommonResponse{
		Ok:  true,
		Msg: "创建成功",
	}, nil
}
