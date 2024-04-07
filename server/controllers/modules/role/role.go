package roleControllersModules

import (
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/services"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

var RoleController = &roleController{}

type roleController struct {
}

// CreateRole 创建角色
// @Summary 创建角色
// @Schemes
// @Description 创建角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /roles [post]
func (r *roleController) CreateRole(context *gin.Context) {
	var role dto.CreateRoleParams
	err := context.ShouldBind(&role)
	if err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	if err = services.RoleService.CreateRole(&role); err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, "创建成功")
}

// DeleteRole 删除角色
// @Summary 删除角色
// @Schemes
// @Description 删除角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /roles/{id} [delete]
func (r *roleController) DeleteRole(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterTypeError(context, "id不能为空")
		return
	}

	err := services.RoleService.DeleteRole(id)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}

	utils.Response.Success(context, "删除成功")
}

// UpdateRole 更新角色
// @Summary 更新角色
// @Schemes
// @Description 更新角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /roles/{id} [patch]
func (r *roleController) UpdateRole(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterTypeError(context, "id不能为空")
		return
	}
	var role dto.UpdateRoleParams
	err := context.ShouldBind(&role)
	if err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	if err = services.RoleService.UpdateRole(&role, id); err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, "更新成功")
}

// GetRoles 获取角色
// @Summary 获取角色
// @Schemes
// @Description 获取角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /roles [get]
func (r *roleController) GetRoles(context *gin.Context) {
	roles, err := services.RoleService.GetRoles()
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, roles)
}
