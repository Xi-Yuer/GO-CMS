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
	user, err := services.UserService.GetUser(id)
	if err != nil {
		utils.Log.Error(err)
	}
	utils.Response.Success(context, user)
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
		utils.Response.ParameterTypeError(context, err)
		return
	}
	err = services.UserService.CreateUser(&user)
	if err != nil {
		utils.Response.ServerError(context, err)
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

}
