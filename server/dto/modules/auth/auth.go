package authResponsiesModules

import usersResponsiesModules "github.com/Xi-Yuer/cms/dto/modules/users"

type LoginRequestParams struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
	Captcha  string `form:"captcha" binding:"required"`
}

type LoginResponse struct {
	Token string                                     `json:"token"`
	User  *usersResponsiesModules.SingleUserResponse `json:"user"`
}

type CreateUserRoleRecordParams struct {
	RoleID []string `form:"roleID" json:"roleID"`
}

type CreateRolePermissionRecordParams struct {
	RoleID      string   `form:"roleID" json:"roleID"`
	PageID      []string `form:"pageID" json:"pageID"`
	InterfaceID []string `form:"interfaceID" json:"interfaceID"`
}
