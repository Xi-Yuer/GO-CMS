package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	roleModlel "micro/model/role"
	"micro/rpc/role/internal/svc"
	"micro/rpc/role/roleRPC"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *roleRPC.UpdateRoleRequest) (*roleRPC.CommonResponse, error) {
	existLogic := NewRoleIDsHasBeenExistLogic(l.ctx, l.svcCtx)
	exist, err := existLogic.RoleIDsHasBeenExist(&roleRPC.RoleIDsHasBeenExistRequest{Ids: []string{in.Id}})

	if err != nil {
		return nil, err
	}

	if !exist.Ok {
		return nil, errors.New("角色ID不存在")
	}

	if err = l.svcCtx.GormDB.Model(&roleModlel.Role{}).Where("id = ?", in.Id).Updates(&roleModlel.Role{
		RoleName:    in.RoleName,
		Description: in.Description,
		CanEdit:     in.CanEdit,
	}).Error; err != nil {
		return nil, err
	}

	return &roleRPC.CommonResponse{
		Ok:  true,
		Msg: "更新成功",
	}, nil
}
