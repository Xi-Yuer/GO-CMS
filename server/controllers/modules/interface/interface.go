package interfaceControllerModules

import (
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/services"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

var InterfaceController = &interfaceControllerModules{}

type interfaceControllerModules struct{}

// CreateInterface 创建接口
// @Summary 创建接口
// @Description 创建接口
// @Tags 接口管理
// @Accept json
// @Produce json
// @Router /interface [post]
func (i *interfaceControllerModules) CreateInterface(context *gin.Context) {
	var params dto.CreateInterfaceRequest
	err := context.ShouldBind(&params)
	if err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}

	if err := services.InterfaceService.CreateInterface(&params); err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, "创建成功")
}

// UpdateInterface 更新接口
// @Summary 更新接口
// @Description 更新接口
// @Tags 接口管理
// @Accept json
// @Produce json
// @Router /interface/{id} [patch]
func (i *interfaceControllerModules) UpdateInterface(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterMissing(context, "id不能为空")
		return
	}
	var params dto.UpdateInterfaceRequest

	if err := context.ShouldBind(&params); err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}

	if err := services.InterfaceService.UpdateInterfaceByID(id, &params); err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, "更新成功")
}

// GetInterfaceByPageID 获取接口（根据页面ID）
// @Summary 获取接口（根据页面ID）
// @Description 获取接口（根据页面ID）
// @Tags 接口管理
// @Accept json
// @Produce json
// @Router /interface/page/{id} [get]
func (i *interfaceControllerModules) GetInterfaceByPageID(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterMissing(context, "id不能为空")
		return
	}
	interfaceResponses := services.InterfaceService.GetInterfaceByPageID(id)
	utils.Response.Success(context, interfaceResponses)
}

// DeleteInterface 删除接口
// @Summary 删除接口
// @Description 删除接口
// @Tags 接口管理
// @Accept json
// @Produce json
// @Router /interface/{id} [delete]
func (i *interfaceControllerModules) DeleteInterface(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterMissing(context, "id不能为空")
		return
	}

	if err := services.InterfaceService.DeleteInterfaceByID(id); err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, "删除成功")
}

// GetInterfacesByRoleID 获取接口（根据角色ID）
// @Summary 获取接口（根据角色ID）
// @Description 获取接口（根据角色ID）
// @Tags 接口管理
// @Accept json
// @Produce json
// @Router /interface/role/{id} [get]
func (i *interfaceControllerModules) GetInterfacesByRoleID(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterMissing(context, "id不能为空")
		return
	}
	interfaceResponses, err := services.RolesAndInterfacesService.GetInterfacesByRoleID(id)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, interfaceResponses)
}

// GetAllInterface 获取接口（所有页面）
// @Summary 获取接口（所有页面）
// @Description 获取接口（所有页面）
// @Tags 接口管理
// @Accept json
// @Produce json
// @Router /interface [get]
func (i *interfaceControllerModules) GetAllInterface(context *gin.Context) {
	allInterface, err := services.InterfaceService.GetAllInterface()
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, allInterface)
}
