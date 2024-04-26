package commonResponsiesModules

type ExportExcelResponse struct {
	IDs []string `json:"ids" form:"ids" binding:"required"`
}

type HasTotalResponseData struct {
	Total int `json:"total"`
	List  any `json:"list"`
}
