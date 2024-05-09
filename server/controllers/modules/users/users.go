package userControllersModules

import (
	"github.com/Xi-Yuer/cms/constant"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/services"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

var UserController = &userController{}

type userController struct{}

// CreateUser 创建用户
// @Summary 创建新用户
// @Description 创建新用户并将其添加到系统中
// @Tags 用户管理
// @Accept json
// @Produce json
// @Router /users [post]
func (u *userController) CreateUser(context *gin.Context) {
	var user dto.CreateSingleUserRequest
	err := context.ShouldBind(&user)
	if err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	err = services.UserService.CreateUser(&user)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, nil)
}

// GetUser 获取用户
// @Summary 根据ID获取用户信息
// @Schemes
// @Description 根据ID获取用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /users/{id} [get]
func (u *userController) GetUser(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterTypeError(context, "id不能为空")
		return
	}
	user, err := services.UserService.GetUser(id)
	if err != nil {
		utils.Response.NotFound(context, "资源不存在")
		return
	}
	utils.Response.Success(context, user)
}

// GetUsers 查询用户
// @Summary 查询用户
// @Description 根据查询参数查询用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Body
// @Router /users/query [post]
func (u *userController) GetUsers(context *gin.Context) {
	var params dto.QueryUsersParams
	err := context.ShouldBind(&params)
	if err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	if params.Limit > 100 || params.Limit < 0 {
		utils.Response.ParameterTypeError(context, "limit参数不正确")
		return
	}
	users, err := services.UserService.FindUserByParams(&params)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, users)
}

// FindUserByParamsAndOutRoleID 查询角色以外的用户
// @Summary 查询角色以外的用户
// @Description 查询角色以外的用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Body
// @Router /users/query/role/:id [post]
func (u *userController) FindUserByParamsAndOutRoleID(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterTypeError(context, "id不能为空")
		return
	}
	var params dto.QueryUsersParams
	err := context.ShouldBind(&params)
	if err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	if params.Limit > 100 || params.Limit < 0 {
		utils.Response.ParameterTypeError(context, "limit参数不正确")
		return
	}
	users, err := services.UserService.FindUserByParamsAndOutRoleID(id, &params)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, users)
}

// UpdateUser 更新用户
// @Summary 更新用户信息
// @Description 更新现有用户的信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Router /users/{id} [patch]
func (u *userController) UpdateUser(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterTypeError(context, "id不能为空")
		return
	}
	var user dto.UpdateUserRequest
	err := context.ShouldBind(&user)
	if err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	jwtPayload, exist := context.Get(constant.JWTPAYLOAD)
	// 判断是否为管理员或者用户自己
	if jwtPayload.(*dto.JWTPayload).IsAdmin == 1 || jwtPayload.(*dto.JWTPayload).ID == id {
		err = services.UserService.UpdateUser(&user, id)
		if err != nil {
			utils.Response.ServerError(context, err.Error())
			return
		}
		utils.Response.Success(context, nil)
		return
	}
	// 判断是否为普通用户
	if !exist || jwtPayload.(*dto.JWTPayload).ID != id {
		utils.Response.NoPermission(context, "暂无权限")
		return
	}
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 根据用户ID删除用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Router /users/{id} [delete]
func (u *userController) DeleteUser(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterTypeError(context, "id不能为空")
		return
	}
	err := services.UserService.DeleteUser(id)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, nil)
}

// GetUserByRoleID 查询用户（查询某个角色下的所有用户）
// @Summary 查询用户（查询某个角色下的所有用户）
// @Description 查询用户（查询某个角色下的所有用户）
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Router /users/role/{id} [get]
func (u *userController) GetUserByRoleID(context *gin.Context) {
	roleID := context.Param("id")
	var params dto.Page
	if roleID == "" {
		utils.Response.ParameterTypeError(context, "角色ID不能为空")
		return
	}
	err := context.ShouldBind(&params)
	if err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	if *params.Limit > 100 {
		utils.Response.ParameterTypeError(context, "limit不能大于100")
		return
	}
	singleResponses, err := services.UserService.GetUserByRoleID(roleID, params)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, singleResponses)
}

// ExportExcel 导出用户
// @Summary 导出用户
// @Description 导出用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Router /users/export [post]
func (u *userController) ExportExcel(context *gin.Context) {
	var IDs dto.ExportExcelResponse
	err := context.ShouldBind(&IDs)
	if err != nil {
		utils.Response.ParameterMissing(context, err.Error())
		return
	}
	responses, err := services.UserService.ExportExcel(&IDs)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	var data [][]interface{}
	data = append(data, []interface{}{"ID", "账号", "昵称", "头像", "状态", "部门ID", "角色ID", "接口权限", "创建时间", "更新时间"})
	for _, response := range responses {
		data = append(data, []interface{}{response.ID, response.Account, response.Nickname, response.Avatar, response.Status, response.DepartmentID, response.RolesID, response.InterfaceDic, response.CreateTime.Format("2006/01/02 03:04:05"), response.UpdateTime.Format("2006/01/02 03:04:05")})
	}
	if err := utils.ExportExcel(context, data, "用户表"); err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
}
