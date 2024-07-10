package logic

import (
	"context"
	userModel "micro/model/user"

	"micro/rpc/user/internal/svc"
	"micro/rpc/user/userRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *userRPC.UpdateUserRequest) (*userRPC.CommonResponse, error) {
	existLogic := NewUserIDHasBeenExistLogic(l.ctx, l.svcCtx)
	exist, err := existLogic.UserIDHasBeenExist(&userRPC.DeleteUserRequest{Id: in.Id})

	if err != nil {
		return &userRPC.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		}, err
	}
	if !exist.Ok {
		return &userRPC.CommonResponse{
			Ok:  false,
			Msg: "用户不存在",
		}, nil
	}

	if err := l.svcCtx.GormDB.Model(&userModel.User{}).Where("id", in.Id).Updates(&userModel.User{
		Password:     in.Password,
		NickName:     in.Nickname,
		Avatar:       in.Avatar,
		Status:       in.Status,
		DepartmentID: in.Department,
		IsAdmin:      in.IsAdmin,
	}).Error; err != nil {
		return &userRPC.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		}, nil
	}
	return &userRPC.CommonResponse{
		Ok:  true,
		Msg: "更新成功",
	}, nil
}
