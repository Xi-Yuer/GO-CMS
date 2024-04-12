package rolesResponsiesModules

type CreateRoleParams struct {
	RoleName    string   `json:"roleName" form:"roleName" binding:"required"`
	Description string   `json:"description" form:"description" binding:"required"`
	PageID      []string `form:"pageID" json:"pageID"`
	InterfaceID []string `form:"interfaceID" json:"interfaceID"`
}

type UpdateRoleParams struct {
	RoleName    string   `json:"roleName" form:"roleName"`
	Description string   `json:"description" form:"description"`
	PageID      []string `json:"pageID" form:"pageID"`
	InterfaceID []string `form:"interfaceID" json:"interfaceID"`
}

type QueryRoleListParams struct {
	ID          string `form:"id"`
	Limit       int    `form:"limit"`
	Offset      int    `form:"offset"`
	RoleName    string `form:"roleName"`
	Description string `form:"description"`
	StartTime   string `form:"startTime"`
	EndTime     string `form:"endTime"`
}

type SingleRoleResponse struct {
	ID           string   `json:"id"`
	RoleName     string   `json:"roleName"`
	Description  string   `json:"description"`
	PagesID      []string `json:"pagesID"`
	InterfacesID []string `json:"interfacesID"`
	CreateTime   string   `json:"createTime"`
	UpdateTime   string   `json:"updateTime"`
}
