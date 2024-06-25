package logic

import (
	"context"
	roleModlel "micro/model/role"

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

func (l *RoleNameHasBeenExistLogic) RoleNameHasBeenExist(in *roleRPC.RoleNamesHasBeenExistRequest) (*roleRPC.CommonResponse, error) {
	count := int64(0)
	if err := l.svcCtx.GormDB.Where("role_name = ?", in.Names).First(&roleModlel.Role{}).Count(&count).Error; err != nil {
		return &roleRPC.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		}, nil
	}
	if count > 0 {
		return &roleRPC.CommonResponse{
			Ok:  true,
			Msg: "角色名已存在",
		}, nil
	}
	return &roleRPC.CommonResponse{
		Ok:  false,
		Msg: "",
	}, nil
}
