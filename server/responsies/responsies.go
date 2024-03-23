package responsies

import usersResponsiesModules "github.com/Xi-Yuer/cms/responsies/modules/users"

// UsersSingleResponse 查询单个用户
type UsersSingleResponse usersResponsiesModules.SingleUserResponse

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
