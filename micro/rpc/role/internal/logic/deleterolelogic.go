package logic

import (
	"context"
	"errors"
	roleModlel "micro/model/role"

	"micro/rpc/role/internal/svc"
	"micro/rpc/role/roleRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(in *roleRPC.DeleteRoleRequest) (*roleRPC.CommonResponse, error) {
	existLogic := NewRoleIDsHasBeenExistLogic(l.ctx, l.svcCtx)
	exist, err := existLogic.RoleIDsHasBeenExist(&roleRPC.RoleIDsHasBeenExistRequest{Ids: []string{in.Id}})

	if err != nil {
		return nil, err
	}

	if !exist.Ok {
		return nil, errors.New("角色ID不存在")
	}

	if err := l.svcCtx.GormDB.Delete(&roleModlel.Role{}, "id = ?", in.Id).Error; err != nil {
		return &roleRPC.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		}, nil
	}
	return &roleRPC.CommonResponse{
		Ok:  true,
		Msg: "删除成功",
	}, nil
}
