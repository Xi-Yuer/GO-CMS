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
	if err := l.svcCtx.GormDB.Where("id = ?", in.Id).Unscoped().Delete(&userModel.User{}).Error; err != nil {
		return &userRPC.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		}, nil
	}
	return &userRPC.CommonResponse{
		Ok:  true,
		Msg: "删除成功",
	}, nil
}
