package authMiddleWareModule

import (
	"github.com/Xi-Yuer/cms/constant"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(context *gin.Context) {
	// 白名单
	whiteList := []string{"/api/v1/auth/login", "/api/v1/auth/captcha"}
	for _, v := range whiteList {
		if v == context.Request.URL.Path {
			context.Next()
			return
		}
	}
	// 进行身份验证的逻辑
	token := context.GetHeader(constant.AUTHORIZATION)
	tokenHs256, err := utils.Jsonwebtoken.ParseTokenHs256(token)
	if err != nil {
		utils.Response.NoAuth(context, "Token验证失败")
		context.Abort()
		return
	}
	// 验证通过后，将用户信息存储到context中，方便后续使用
	context.Set(constant.JWTPAYLOAD, tokenHs256)
	context.Next()
}
