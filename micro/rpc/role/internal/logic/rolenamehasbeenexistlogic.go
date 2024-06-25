package logic

import (
	"context"

	"micro/rpc/role/internal/svc"
	"micro/rpc/role/roleRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleNameHasBeenExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleNameHasBeenExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleNameHasBeenExistLogic {
	return &RoleNameHasBeenExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleNameHasBeenExistLogic) RoleNameHasBeenExist(in *roleRPC.DeleteRoleRequest) (*roleRPC.CommonResponse, error) {
	// todo: add your logic here and delete this line

	return &roleRPC.CommonResponse{}, nil
}
