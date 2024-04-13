package dto

import (
	authResponsiesModules "github.com/Xi-Yuer/cms/dto/modules/auth"
	commitsResponsiesModules "github.com/Xi-Yuer/cms/dto/modules/commits"
	departmentResponsiesModules "github.com/Xi-Yuer/cms/dto/modules/department"
	interfaceResponsiesModules "github.com/Xi-Yuer/cms/dto/modules/interface"
	logsResponsiesModules "github.com/Xi-Yuer/cms/dto/modules/logs"
	pagesResponsiesModules "github.com/Xi-Yuer/cms/dto/modules/pages"
	rolesResponsiesModules "github.com/Xi-Yuer/cms/dto/modules/roles"
	usersResponsiesModules "github.com/Xi-Yuer/cms/dto/modules/users"
)

// LoginResponse 用户登录
type LoginResponse authResponsiesModules.LoginResponse

// UsersSingleResponse 查询单个用户
type UsersSingleResponse usersResponsiesModules.SingleUserResponse

// SingleUserByRoleIDResponse 根据角色id查询单个用户
type SingleUserByRoleIDResponse usersResponsiesModules.SingleUserByRoleIDResponse

type SingleUserResponseHasPassword usersResponsiesModules.SingleUserResponseHasPassword

// CreateSingleUserRequest 创建用户
type CreateSingleUserRequest usersResponsiesModules.CreateUserParams

// UpdateSingleUserRequest 更新用户
type UpdateSingleUserRequest usersResponsiesModules.UpdateUserParams

// QueryUsersParams 查询用户
type QueryUsersParams usersResponsiesModules.QueryUsersParams

// UpdateUserRequest 更新用户
type UpdateUserRequest usersResponsiesModules.UpdateUserParams

// Page 查询参数
type Page usersResponsiesModules.Page

// LoginRequestParams 登录请求
type LoginRequestParams = authResponsiesModules.LoginRequestParams

// JWTPayload 生成JWT
type JWTPayload = usersResponsiesModules.JWTPayload

// CreateRoleParams 创建角色
type CreateRoleParams = rolesResponsiesModules.CreateRoleParams

// UpdateRoleParams 更新角色
type UpdateRoleParams = rolesResponsiesModules.UpdateRoleParams

// QueryRolesParams 查询角色
type QueryRolesParams = rolesResponsiesModules.QueryRoleListParams

// SingleRoleResponse 查询单个角色
type SingleRoleResponse = rolesResponsiesModules.SingleRoleResponse

// AuthorizationManagementParams 给用户分配角色
type AuthorizationManagementParams = authResponsiesModules.CreateUserRoleRecordParams

// CreatePageParams 创建页面
type CreatePageParams = pagesResponsiesModules.CreatePageParams

// UpdatePageRequest 更新页面
type UpdatePageRequest = pagesResponsiesModules.UpdatePageRequest

// SinglePageResponse 查询页面
type SinglePageResponse = pagesResponsiesModules.SinglePageResponse

// CreateRolePermissionRecordParams 分配角色权限
type CreateRolePermissionRecordParams = authResponsiesModules.CreateRolePermissionRecordParams

// CreateDepartmentRequest 创建部门
type CreateDepartmentRequest = departmentResponsiesModules.CreateDepartmentRequest

// DepartmentResponse 查询部门
type DepartmentResponse = departmentResponsiesModules.DepartmentResponse

// UpdateDepartmentRequest 更新部门
type UpdateDepartmentRequest = departmentResponsiesModules.UpdateDepartmentRequest

// CreateInterfaceRequest 创建接口
type CreateInterfaceRequest = interfaceResponsiesModules.CreateInterfaceRequest

// GetInterfaceResponse 获取接口
type GetInterfaceResponse = interfaceResponsiesModules.GetInterfaceResponse

// UpdateInterfaceRequest 更新接口
type UpdateInterfaceRequest = interfaceResponsiesModules.UpdateInterfaceRequest

// CreateLogRecordRequest 创建日志
type CreateLogRecordRequest = logsResponsiesModules.CreateLogRecordRequest

// GetLogRecordResponse 获取日志
type GetLogRecordResponse = logsResponsiesModules.GetLogRecordResponse

// CommitResponse 提交记录
type CommitResponse = commitsResponsiesModules.CommitResponse
