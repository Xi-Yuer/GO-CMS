package logic

import (
	"context"
	"errors"
	roleModlel "micro/model/role"
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
	var roles []roleModlel.Role
	if err := l.svcCtx.GormDB.Where("id in (?)", in.Ids).Find(&roles).Error; err != nil {
		return nil, err
	}

	if len(roles) != len(in.Ids) {
		return nil, errors.New("角色ID不存在")
	}

	return &roleRPC.CommonResponse{
		Ok:  true,
		Msg: "",
	}, nil

}
