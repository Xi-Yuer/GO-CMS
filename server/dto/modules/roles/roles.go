package rolesResponsiesModules

type CreateRoleParams struct {
	RoleName    string   `json:"roleName" form:"roleName" binding:"required"`
	Description string   `json:"description" form:"description" binding:"required"`
	PageID      []string `form:"pageID" json:"pageID"`
}

type UpdateRoleParams struct {
	RoleName    string   `json:"roleName" form:"roleName"`
	Description string   `json:"description" form:"description"`
	PageID      []string `json:"pageID" form:"pageID"`
}

type QueryRoleListParams struct{}

type SingleRoleResponse struct {
	ID          string   `json:"id"`
	RoleName    string   `json:"roleName"`
	Description string   `json:"description"`
	PagesID     []string `json:"pagesID"`
	CreateTime  string   `json:"createTime"`
	UpdateTime  string   `json:"updateTime"`
}
