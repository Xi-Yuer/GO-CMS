package templateResponsiesModules

type CreateTemplateRequestParams struct {
	TableName string    `form:"tableName" binding:"required"`
	Fields    *[]Fields `form:"fields" binding:"required"`
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
	DTOFile        Code `json:"dtoFile"`
}
type Web struct {
}

type Code struct {
	Code string `json:"code"`
	Lang string `json:"lang"`
}
