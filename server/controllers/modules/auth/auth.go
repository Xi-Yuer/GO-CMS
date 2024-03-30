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
// @Tags 权限
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
	err, s := services.AuthService.Login(&loginParams)
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
// @Tags 权限
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/captcha [get]
func (a *authController) Captcha(context *gin.Context) {
	utils.Captcha.GenerateCaptcha(context, 4)
}

// AuthorizationManagementController 给用户分配角色
// @Summary 给用户分配角色
// @Schemes
// @Description 给用户分配角色
// @Tags 权限
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/authorization [post]
func (a *authController) AuthorizationManagementController(context *gin.Context) {
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
	err = services.AuthService.AuthorizationManagement(user.(*dto.JWTPayload).ID, &authorizationManagementParams)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, nil)
}
