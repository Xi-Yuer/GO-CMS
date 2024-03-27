package userControllersModules

import (
	"github.com/Xi-Yuer/cms/responsies"
	"github.com/Xi-Yuer/cms/services"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

var UserController = &userController{}

type userController struct{}

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
		utils.Response.NotFound(context, nil)
		return
	}
	utils.Response.Success(context, user)
}

// FindUserByParams 查询用户
// @Summary 查询用户
// @Description 根据查询参数查询用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Body
// @Router /users [get]
func (u *userController) FindUserByParams(content *gin.Context) {
	var params responsies.QueryUsersParams
	err := content.ShouldBind(&params)
	if err != nil {
		utils.Response.ParameterTypeError(content, err.Error())
		return
	}
	users, err := services.UserService.FindUserByParams(&params)
	if err != nil {
		utils.Response.ServerError(content, err.Error())
		return
	}
	utils.Response.Success(content, users)
}

// CreateUser 创建用户
// @Summary 创建新用户
// @Description 创建新用户并将其添加到系统中
// @Tags 用户管理
// @Accept json
// @Produce json
// @Router /users [post]
func (u *userController) CreateUser(context *gin.Context) {
	var user responsies.CreateSingleUserRequest
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

// UpdateUser 更新用户
// @Summary 更新用户信息
// @Description 更新现有用户的信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Router /users/{id} [put]
func (u *userController) UpdateUser(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterTypeError(context, "id不能为空")
		return
	}
	var user responsies.UpdateUserRequest
	err := context.ShouldBind(&user)
	if err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	err = services.UserService.UpdateUser(&user, id)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, nil)
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

// GetUsers 获取用户列表
// @Summary 获取用户列表
// @Description 获取系统中所有用户的列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param page query int false "页码，默认为1"
// @Param limit query int false "每页显示的用户数量，默认为10"
// @Router /users [get]
func (u *userController) GetUsers(context *gin.Context) {
	var Page responsies.Page
	err := context.ShouldBind(&Page)
	if err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	users, err := services.UserService.GetUsers(Page)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, users)
}
