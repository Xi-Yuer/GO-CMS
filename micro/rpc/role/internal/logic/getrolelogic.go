package logic

import (
	"context"
	"errors"
	roleModlel "micro/model/role"

	"micro/rpc/role/internal/svc"
	"micro/rpc/role/roleRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleLogic) GetRole(in *roleRPC.GetRoleRequest) (*roleRPC.GetRoleResponse, error) {
	existLogic := NewRoleIDsHasBeenExistLogic(l.ctx, l.svcCtx)
	exist, err := existLogic.RoleIDsHasBeenExist(&roleRPC.RoleIDsHasBeenExistRequest{Ids: []string{in.Id}})

	if err != nil {
		return nil, err
	}

	if !exist.Ok {
		return nil, errors.New("角色ID不存在")
	}

	role := roleModlel.Role{}
	err = l.svcCtx.GormDB.Model(&roleModlel.Role{}).First(&role, in.Id).Error
	if err != nil {
		return nil, err
	}
	return &roleRPC.GetRoleResponse{
		Id:          role.ID,
		RoleName:    role.RoleName,
		Description: role.Description,
		CanEdit:     role.CanEdit,
		CreateTime:  role.CreatedAt.String(),
		UpdateTime:  role.UpdatedAt.String(),
	}, nil
}
