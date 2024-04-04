package authControllersModules

import (
	"github.com/Xi-Yuer/cms/constant"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/services"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

var AuthController = &authController{}

type authController struct{}

// Login 登录
// @Summary 登录
// @Schemes
// @Description 登录
// @Tags 权限管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/login [post]
func (a *authController) Login(context *gin.Context) {
	var loginParams dto.LoginRequestParams
	if err := context.ShouldBind(&loginParams); err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	// 校验验证码
	verifyCaptcha := utils.Captcha.VerifyCaptcha(context, loginParams.Captcha)
	if !verifyCaptcha {
		utils.Response.ParameterTypeError(context, "验证码有误")
		return
	}

	// 登录
	s, err := services.AuthService.Login(&loginParams)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, s)
}

// Captcha 获取验证码
// @Summary 获取验证码
// @Schemes
// @Description 获取验证码
// @Tags 权限管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/captcha [get]
func (a *authController) Captcha(context *gin.Context) {
	utils.Captcha.GenerateCaptcha(context, 4)
}

// CreateUserRoleRecordController 给用户分配角色
// @Summary 给用户分配角色
// @Schemes
// @Description 给用户分配角色
// @Tags 权限管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/bindRoles [post]
func (a *authController) CreateUserRoleRecordController(context *gin.Context) {
	var authorizationManagementParams dto.AuthorizationManagementParams
	err := context.ShouldBind(&authorizationManagementParams)
	if err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	user, exists := context.Get(constant.JWTPAYLOAD)
	if !exists {
		utils.Response.NoPermission(context, "暂无权限")
	}
	err = services.AuthService.CreateUserRoleRecord(user.(*dto.JWTPayload).ID, &authorizationManagementParams)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, nil)
}

// CreateRolePermissionRecordController 给角色分配权限
// @Summary 给角色分配权限
// @Schemes
// @Description 给角色分配权限
// @Tags 权限管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/bindPermissions [post]
func (a *authController) CreateRolePermissionRecordController(context *gin.Context) {
	var createRolePermissionRecordParams dto.CreateRolePermissionRecordParams
	err := context.ShouldBind(&createRolePermissionRecordParams)
	if err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	if err = services.AuthService.CreateRolePermissionsRecord(&createRolePermissionRecordParams); err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, nil)
}

// GetUserMenus 获取用户菜单
// @Summary 获取用户菜单
// @Description 获取用户菜单
// @Tags 权限管理
// @Accept json
// @Produce json
// @Router /users [get]
func (a *authController) GetUserMenus(context *gin.Context) {
	jwtPayload, exist := context.Get(constant.JWTPAYLOAD)
	if !exist {
		utils.Response.ParameterMissing(context, "用户id不能为空")
		return
	}
	menus, err := services.AuthService.GetUserMenus(jwtPayload.(*dto.JWTPayload).ID)
	if err != nil {
		return
	}
	utils.Response.Success(context, menus)
}
