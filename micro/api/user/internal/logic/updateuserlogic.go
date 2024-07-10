package logic

import (
	"context"
	"micro/api/user/internal/svc"
	"micro/api/user/internal/types"
	"micro/rpc/user/userRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserRequest) (resp *types.UpdateUserResponse, err error) {
	response, _ := l.svcCtx.UserService.UpdateUser(l.ctx, &userRPC.UpdateUserRequest{
		Id:         req.ID,
		Password:   req.Password,
		Nickname:   req.Nickname,
		Avatar:     req.Avatar,
		Status:     req.Status,
		Department: req.DepartmentID,
		IsAdmin:    req.IsAdmin,
	})
	if response.Ok {
		return &types.UpdateUserResponse{
			Code: 200,
			Msg:  response.Msg,
		}, nil
	}
	return &types.UpdateUserResponse{
		Code: 500,
		Msg:  response.Msg,
	}, nil
}
