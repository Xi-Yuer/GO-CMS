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
	ID          string `json:"id"`
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
	CanEdit     int
	PageID      []string `json:"pageID"`
	InterfaceID []string `json:"interfaceID"`
	CreateTime  string   `json:"createTime"`
	UpdateTime  string   `json:"updateTime"`
}

type CreateOneRecord struct {
	UserID string `form:"userID"`
	RoleID string `form:"roleID"`
}

type DeleteOneRecord struct {
	UserID string `form:"userID"`
	RoleID string `form:"roleID"`
}
