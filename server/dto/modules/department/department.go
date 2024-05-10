package departmentResponsiesModules

type CreateDepartmentRequest struct {
	DepartmentName        string  `form:"departmentName" binding:"required"`
	ParentDepartment      *string `form:"parentDepartment"`
	DepartmentDescription string  `form:"departmentDescription"`
	DepartmentOrder       int     `form:"departmentOrder"`
}

type DepartmentResponse struct {
	ID                    string                `json:"id"`
	DepartmentName        string                `json:"departmentName"`
	ParentDepartment      *string               `json:"parentDepartment"`
	DepartmentOrder       int                   `json:"departmentOrder"`
	DepartmentDescription *string               `json:"departmentDescription"`
	Children              []*DepartmentResponse `json:"children"`
	CreateTime            string                `json:"createTime"`
	UpdateTime            string                `json:"updateTime"`
}

type UpdateDepartmentRequest struct {
	DepartmentName        string `form:"departmentName"`
	ParentDepartment      string `form:"parentDepartment"`
	DepartmentDescription string `json:"departmentDescription"`
	DepartmentOrder       int    `form:"departmentOrder"`
}
