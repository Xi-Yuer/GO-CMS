package authResponsiesModules

type LoginRequestParams struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
	Captcha  string `form:"captcha" binding:"required"`
}
