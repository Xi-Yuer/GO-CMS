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
	Controller string `json:"controller"`
	Service    string `json:"service"`
	Repository string `json:"repository"`
	DTO        string `json:"dto"`
}
