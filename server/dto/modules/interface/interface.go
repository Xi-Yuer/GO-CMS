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
	CreateTime      string `json:"createTime"`
	UpdateTime      string `json:"updateTime"`
	InterfaceDic    string `json:"interfaceDic"`
}

type UpdateInterfaceRequest struct {
	InterfaceName   string `form:"interfaceName" binding:"required"`
	InterfaceMethod string `form:"interfaceMethod" binding:"required"`
	InterfacePageID string `form:"interfacePageID" binding:"required"`
	InterfacePath   string `form:"interfacePath" binding:"required"`
	InterfaceDesc   string `form:"interfaceDesc" binding:"required"`
	InterfaceDic    string `form:"interfaceDic" binding:"required"`
}
