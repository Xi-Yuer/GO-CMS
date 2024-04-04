package authResponsiesModules

type LoginRequestParams struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
	Captcha  string `form:"captcha" binding:"required"`
}

type CreateUserRoleRecordParams struct {
	RoleID []string `form:"roleID" json:"roleID"`
}

type CreateRolePermissionRecordParams struct {
	RoleID string   `form:"roleID" json:"roleID" binding:"required"`
	PageID []string `form:"pageID" json:"pageID"`
}
