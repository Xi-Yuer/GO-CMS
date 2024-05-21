package templateResponsiesModules

type CreateTemplateRequestParams struct {
	Package   string    `form:"package" binding:"required"`
	TableName string    `form:"tableName" binding:"required"`
	Fields    *[]Fields `form:"fields" binding:"required"`
}

type DownloadTemplateRequestParams struct {
	TableName  string `form:"tableName" binding:"required"`
	Controller string `form:"controller" binding:"required"`
	Service    string `form:"service" binding:"required"`
	Repository string `form:"repository" binding:"required"`
	Route      string `form:"route" binding:"required"`
	DTO        string `form:"dto" binding:"required"`
	SearchForm string `form:"searchForm" binding:"required"`
	Table      string `form:"table" binding:"required"`
	TableHook  string `form:"tableHook" binding:"required"`
}

type Fields struct {
	Name    string `form:"name" binding:"required"`
	Type    string `form:"type" binding:"required"`
	Default string `form:"default" binding:"required"`
}

type CreateTemplateResponse struct {
	Server Server `json:"server"`
	Web    Web    `json:"web"`
}

type Server struct {
	ControllerFile Code `json:"controllerFile"`
	ServiceFile    Code `json:"serviceFile"`
	RepositoryFile Code `json:"repositoryFile"`
	RouteFile      Code `json:"routeFile"`
	DTOFile        Code `json:"dtoFile"`
}
type Web struct {
	React React `json:"react"`
}

type React struct {
	SearchForm Code `json:"searchForm"`
	Table      Code `json:"table"`
	TableHook  Code `json:"tableHook"`
}

type Code struct {
	Code string `json:"code"`
	Lang string `json:"lang"`
}
