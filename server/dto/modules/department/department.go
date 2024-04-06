package departmentResponsiesModules

type CreateDepartmentRequest struct {
	DepartmentName   string `form:"departmentName" binding:"required"`
	ParentDepartment string `form:"parentDepartment"`
	DepartmentOrder  int    `form:"departmentOrder"`
}

type DepartmentResponse struct {
	ID               string               `json:"id"`
	DepartmentName   string               `json:"departmentName"`
	ParentDepartment *string              `json:"parentDepartment"`
	DepartmentOrder  int                  `json:"departmentOrder"`
	Children         []DepartmentResponse `json:"children"`
	CreateTime       string               `json:"createTime"`
	UpdateTime       string               `json:"updateTime"`
}
