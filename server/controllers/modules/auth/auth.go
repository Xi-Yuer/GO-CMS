package authControllersModules

import (
	"github.com/Xi-Yuer/cms/config"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/services"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var AuthController = &authController{}

type authController struct{}

// Login 登录
// @Summary 登录
// @Schemes
// @Description 登录
// @Tags 登录
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
// @Tags 登录
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/captcha [get]
func (a *authController) Captcha(context *gin.Context) {
	utils.Captcha.GenerateCaptcha(context, 4)
}

// Cookie 获取一个短期的 Cookie
// @Summary 获取一个短期的 Cookie
// @Schemes
// @Description 获取一个短期的 Cookie 文件下载需要用到
// @Tags 登录
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/cookie [get]
func (a *authController) Cookie(context *gin.Context) {
	userToken := context.Request.Header.Get("Authorization")
	session := sessions.Default(context)
	session.Set("token", userToken)
	_ = session.Save()
	context.SetCookie("token", userToken, 3600, "/", config.Config.APP.DOMAIN, true, true)
	utils.Response.Success(context, "请求成功,可继续下载文件")
}
