package commonResponsiesModules

type ExportExcelResponse struct {
	IDs []string `json:"ids" form:"ids" binding:"required"`
}
