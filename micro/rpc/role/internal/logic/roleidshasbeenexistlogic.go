package logic

import (
	"context"

	"micro/rpc/role/internal/svc"
	"micro/rpc/role/roleRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleIDsHasBeenExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleIDsHasBeenExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleIDsHasBeenExistLogic {
	return &RoleIDsHasBeenExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleIDsHasBeenExistLogic) RoleIDsHasBeenExist(in *roleRPC.RoleIDsHasBeenExistRequest) (*roleRPC.CommonResponse, error) {
	// todo: add your logic here and delete this line

	return &roleRPC.CommonResponse{}, nil
}
