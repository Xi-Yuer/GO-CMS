package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	roleModlel "micro/model/role"
	"micro/rpc/role/internal/svc"
	"micro/rpc/role/roleRPC"
)

type GetRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleListLogic) GetRoleList(in *roleRPC.GetRoleListRequest) (*roleRPC.GetUserListResponse, error) {
	total := int64(0)
	roleList := make([]*roleModlel.Role, 0)
	err := l.svcCtx.GormDB.Find(&roleList).Where(roleModlel.Role{
		ID:          in.Id,
		RoleName:    in.RoleName,
		Description: in.Description,
		CanEdit:     in.CanEdit,
	}).Offset(int(in.Page)).Limit(int(in.PageSize)).Count(&total).Error

	if err != nil {
		return nil, err
	}

	roleListRPC := make([]*roleRPC.GetRoleResponse, 0)

	for _, role := range roleList {
		roleListRPC = append(roleListRPC, &roleRPC.GetRoleResponse{
			Id:          role.ID,
			RoleName:    role.RoleName,
			Description: role.Description,
			CanEdit:     role.CanEdit,
			CreateTime:  role.CreatedAt.String(),
			UpdateTime:  role.UpdatedAt.String(),
		})
	}

	return &roleRPC.GetUserListResponse{
		Total:    int32(total),
		RoleList: roleListRPC,
	}, nil
}
