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
	if err := l.svcCtx.GormDB.Model(&userModel.User{}).Where("id", in.Id).Updates(&userModel.User{
		Password:     in.Password,
		NickName:     in.Nickname,
		Avatar:       in.Avatar,
		Status:       in.Status,
		DepartmentID: in.Department,
		IsAdmin:      in.IsAdmin,
	}).Error; err != nil {
		return nil, err
	}
	return &userRPC.CommonResponse{
		Ok:  true,
		Msg: "更新成功",
	}, nil
}
