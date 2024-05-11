package interfaceResponsiesModules

type CreateInterfaceRequest struct {
	InterfaceName   string `form:"interfaceName" binding:"required"`
	InterfaceMethod string `form:"interfaceMethod" binding:"required"`
	InterfacePageID string `form:"interfacePageID" binding:"required"`
	InterfacePath   string `form:"interfacePath" binding:"required"`
	InterfaceDesc   string `form:"interfaceDesc" binding:"required"`
	InterfaceDic    string `form:"interfaceDic" binding:"required"`
}

type GetInterfaceResponse struct {
	ID              string `json:"id"`
	InterfaceName   string `json:"interfaceName"`
	InterfaceMethod string `json:"interfaceMethod"`
	InterfacePageID string `json:"interfacePageID" `
	InterfacePath   string `json:"interfacePath" `
	InterfaceDesc   string `json:"interfaceDesc"`
	InterfaceDic    string `json:"interfaceDic"`
	CanEdit         bool   `json:"canEdit"`
	CreateTime      string `json:"createTime"`
	UpdateTime      string `json:"updateTime"`
}

type UpdateInterfaceRequest struct {
	InterfaceName   string `form:"interfaceName" binding:"required"`
	InterfaceMethod string `form:"interfaceMethod" binding:"required"`
	InterfacePageID string `form:"interfacePageID" binding:"required"`
	InterfacePath   string `form:"interfacePath" binding:"required"`
	InterfaceDesc   string `form:"interfaceDesc" binding:"required"`
	InterfaceDic    string `form:"interfaceDic" binding:"required"`
}

type AllInterfaceResponse struct {
	Key      string                  `json:"key"`
	Children []*GetInterfaceResponse `json:"children"`
}
