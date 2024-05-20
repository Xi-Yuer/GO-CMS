package pagesResponsiesModules

type CreatePageParams struct {
	PageName      string  `form:"pageName" binding:"required"`
	PagePath      string  `form:"pagePath"`
	PageIcon      string  `form:"pageIcon" binding:"required"`
	PageComponent string  `form:"pageComponent"`
	ParentPage    string  `form:"parentPage"`
	PageOrder     *int    `form:"pageOrder" binding:"required"`
	IsOutSite     bool    `form:"isOutSite"`
	OutSiteLink   *string `form:"outSiteLink"`
}

type SinglePageResponse struct {
	PageID        string                `json:"pageID"`
	PageName      string                `json:"pageName"`
	PagePath      string                `json:"pagePath"`
	PageIcon      string                `json:"pageIcon"`
	PageComponent string                `json:"pageComponent"`
	ParentPage    *string               `json:"parentPage"`
	Children      []*SinglePageResponse `json:"children"`
	PageOrder     int                   `json:"pageOrder"`
	CanEdit       int                   `json:"canEdit"`
	IsOutSite     bool                  `json:"isOutSite"`
	OutSiteLink   *string               `json:"outSiteLink"`
	CreatedTime   string                `json:"createdAt"`
	UpdateTime    string                `json:"updateTime"`
}

type UpdatePageRequest struct {
	PageName      string  `form:"pageName" binding:"required"`
	PagePath      string  `form:"pagePath"`
	PageIcon      string  `form:"pageIcon" binding:"required"`
	PageComponent string  `form:"pageComponent"`
	PageOrder     *int    `form:"pageOrder" binding:"required"`
	IsOutSite     bool    `form:"isOutSite"`
	OutSiteLink   *string `form:"outSiteLink"`
}
