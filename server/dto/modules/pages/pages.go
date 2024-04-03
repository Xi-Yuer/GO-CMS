package pagesResponsiesModules

type CreatePageParams struct {
	PageName      string `form:"pageName" binding:"required"`
	PagePath      string `form:"pagePath" binding:"required"`
	PageIcon      string `form:"pageIcon"`
	PageComponent string `form:"pageComponent" binding:"required"`
	ParentPage    string `form:"parentPage"`
	PageOrder     int    `form:"pageOrder"`
}

type SinglePageResponse struct {
	PageID        string  `json:"pageID"`
	PageName      string  `json:"pageName"`
	PagePath      string  `json:"pagePath"`
	PageIcon      string  `json:"pageIcon"`
	PageComponent string  `json:"pageComponent"`
	ParentPage    *string `json:"parentPage"`
	PageOrder     *int    `json:"pageOrder"`
	CanEdit       int     `json:"canEdit"`
	CreatedTime   string  `json:"createdAt"`
	UpdateTime    string  `json:"updateTime"`
}
