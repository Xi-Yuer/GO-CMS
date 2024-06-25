package logic

import (
	"context"
	"micro/rpc/user/userRPC"
	"micro/shared/snowflake"
	"strconv"

	"micro/api/user/internal/svc"
	"micro/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserRequest) (resp *types.UpdateUserResponse, err error) {
	response, _ := l.svcCtx.UserService.CreateUser(l.ctx, &userRPC.CreateUserRequest{
		Id:         strconv.FormatInt(snowflake.GenID(), 10),
		Account:    req.Account,
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
