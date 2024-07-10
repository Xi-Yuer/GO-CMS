// Code generated by goctl. DO NOT EDIT.
// Source: role.proto

package roleservice

import (
	"context"

	"micro/rpc/role/roleRPC"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommonResponse               = roleRPC.CommonResponse
	CreateRoleRequest            = roleRPC.CreateRoleRequest
	DeleteRoleRequest            = roleRPC.DeleteRoleRequest
	GetRoleListRequest           = roleRPC.GetRoleListRequest
	GetRoleRequest               = roleRPC.GetRoleRequest
	GetRoleResponse              = roleRPC.GetRoleResponse
	GetUserListResponse          = roleRPC.GetUserListResponse
	RoleIDsHasBeenExistRequest   = roleRPC.RoleIDsHasBeenExistRequest
	RoleNamesHasBeenExistRequest = roleRPC.RoleNamesHasBeenExistRequest
	UpdateRoleRequest            = roleRPC.UpdateRoleRequest

	RoleService interface {
		CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CommonResponse, error)
		DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*CommonResponse, error)
		UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*CommonResponse, error)
		GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleResponse, error)
		GetRoleList(ctx context.Context, in *GetRoleListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error)
		RoleNameHasBeenExist(ctx context.Context, in *RoleNamesHasBeenExistRequest, opts ...grpc.CallOption) (*CommonResponse, error)
		RoleIDsHasBeenExist(ctx context.Context, in *RoleIDsHasBeenExistRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	}

	defaultRoleService struct {
		cli zrpc.Client
	}
)

func NewRoleService(cli zrpc.Client) RoleService {
	return &defaultRoleService{
		cli: cli,
	}
}

func (m *defaultRoleService) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	client := roleRPC.NewRoleServiceClient(m.cli.Conn())
	return client.CreateRole(ctx, in, opts...)
}

func (m *defaultRoleService) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	client := roleRPC.NewRoleServiceClient(m.cli.Conn())
	return client.DeleteRole(ctx, in, opts...)
}

func (m *defaultRoleService) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	client := roleRPC.NewRoleServiceClient(m.cli.Conn())
	return client.UpdateRole(ctx, in, opts...)
}

func (m *defaultRoleService) GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*GetRoleResponse, error) {
	client := roleRPC.NewRoleServiceClient(m.cli.Conn())
	return client.GetRole(ctx, in, opts...)
}

func (m *defaultRoleService) GetRoleList(ctx context.Context, in *GetRoleListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error) {
	client := roleRPC.NewRoleServiceClient(m.cli.Conn())
	return client.GetRoleList(ctx, in, opts...)
}

func (m *defaultRoleService) RoleNameHasBeenExist(ctx context.Context, in *RoleNamesHasBeenExistRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	client := roleRPC.NewRoleServiceClient(m.cli.Conn())
	return client.RoleNameHasBeenExist(ctx, in, opts...)
}

func (m *defaultRoleService) RoleIDsHasBeenExist(ctx context.Context, in *RoleIDsHasBeenExistRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	client := roleRPC.NewRoleServiceClient(m.cli.Conn())
	return client.RoleIDsHasBeenExist(ctx, in, opts...)
}
