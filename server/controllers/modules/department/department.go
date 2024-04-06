package departmentControllerModules

import (
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/services"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

var DepartmentController = &departmentController{}

type departmentController struct{}

// CreateDepartment 创建部门
// @Summary 创建部门
// @Description 创建部门
// @Tags 部门管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /department [post]
func (d *departmentController) CreateDepartment(context *gin.Context) {
	var createParams dto.CreateDepartmentRequest
	if err := context.ShouldBind(&createParams); err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	err := services.DepartmentService.CreateDepartment(&createParams)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, nil)
}

// DeleteDepartment 删除部门
// @Summary 删除部门
// @Description 删除部门
// @Tags 部门管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /department/{id} [delete]
func (d *departmentController) DeleteDepartment(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterMissing(context, "id不能为空")
		return
	}
	err := services.DepartmentService.DeleteDepartment(id)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}

	utils.Response.Success(context, nil)
}

// GetDepartments 获取部门
// @Summary 获取部门
// @Description 获取部门
// @Tags 部门管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /department [get]
func (d *departmentController) GetDepartments(context *gin.Context) {
	department, err := services.DepartmentService.GetDepartments()
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, department)
}

// UpdateDepartment 修改部门
// @Summary 修改部门
// @Description 修改部门
// @Tags 部门管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /department/{id} [patch]
func (d *departmentController) UpdateDepartment(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterMissing(context, "id不能为空")
		return
	}
	var params dto.UpdateDepartmentRequest
	err := context.ShouldBind(&params)
	if err != nil {
		utils.Response.ParameterError(context, err.Error())
		return
	}

	if err := services.DepartmentService.UpdateDepartment(id, &params); err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}

	utils.Response.Success(context, nil)
}
