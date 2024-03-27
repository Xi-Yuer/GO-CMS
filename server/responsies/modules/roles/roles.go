package rolesResponsiesModules

type CreateRoleParams struct {
	RoleName    string `json:"roleName" form:"roleName" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}

type UpdateRoleParams struct {
	ID          string `json:"id" form:"id" binding:"required"`
	RoleName    string `json:"roleName" form:"roleName"`
	Description string `json:"description" form:"description"`
}

type QueryRoleListParams struct {
	Offset   *string `json:"offset" form:"offset" binding:"required"`
	Limit    *string `json:"limit" form:"limit" binding:"required"`
	RoleName string  `json:"roleName" form:"roleName"`
}
