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
		utils.Response.NotFound(context, "用户不存在")
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
// @Router /users [get]
func (u *userController) GetUsers(content *gin.Context) {
	var params dto.QueryUsersParams
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
	if !exist || jwtPayload.(*dto.JWTPayload).ID != id {
		utils.Response.NoPermission(context, "暂无权限")
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
