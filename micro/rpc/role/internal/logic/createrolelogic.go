package logic

import (
	"context"
	roleModlel "micro/model/role"
	"micro/rpc/role/internal/svc"
	"micro/rpc/role/roleRPC"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRoleLogic) CreateRole(in *roleRPC.CreateRoleRequest) (*roleRPC.CommonResponse, error) {
	role := &roleModlel.Role{
		ID:          in.Id,
		RoleName:    in.RoleName,
		Description: in.Description,
		CanEdit:     in.CanEdit,
	}
	if err := l.svcCtx.GormDB.Create(role).Error; err != nil {
		return &roleRPC.CommonResponse{
			Ok:  false,
			Msg: err.Error(),
		}, nil
	}

	return &roleRPC.CommonResponse{
		Ok:  true,
		Msg: "创建成功",
	}, nil
}
