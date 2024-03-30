package pagesResponsiesModules

type CreatePageParams struct {
	PageName      string `form:"pageName" binding:"required"`
	PagePath      string `form:"pagePath" binding:"required"`
	PageIcon      string `form:"pageIcon"`
	PageComponent string `form:"pageComponent" binding:"required"`
	ParentPage    string `form:"parentPage"`
}
