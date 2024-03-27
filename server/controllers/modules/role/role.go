package roleControllersModules

import (
	"github.com/Xi-Yuer/cms/responsies"
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
// @Tags 角色
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /roles [post]
func (r *roleController) CreateRole(context *gin.Context) {
	var role responsies.CreateRoleParams
	err := context.ShouldBind(&role)
	if err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
	}
	if err = services.RoleService.CreateRole(&role); err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, "创建成功")
}
