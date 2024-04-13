package authMiddleWareModule

import (
	"github.com/Xi-Yuer/cms/config"
	"github.com/Xi-Yuer/cms/constant"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
	"regexp"
)

func AuthTokenMiddleWare(context *gin.Context) {
	// 白名单
	if pass := WhiteList(context); pass {
		context.Next()
		return
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

func AuthMethodMiddleWare(context *gin.Context) {
	// 白名单
	if pass := WhiteList(context); pass {
		context.Next()
		return
	}

	// 获取请求路径和请求方法
	path := context.Request.URL.Path
	method := context.Request.Method

	// 构建权限键
	permissionKey := method + ":" + path

	// 获取用户权限信息
	user, exist := context.Get(constant.JWTPAYLOAD)
	jwtPayload := &dto.JWTPayload{}
	if exist {
		jwtPayload = user.(*dto.JWTPayload)
	}
	allMethodPermission := jwtPayload.InterfaceDic

	pass := true
	// 遍历权限映射，查找与请求路径匹配的正则表达式
	for reStr, requiredPermission := range constant.PermissionMap {
		re, err := regexp.Compile(reStr)
		if err != nil {
			continue
		}
		// 找到匹配的正则表达式，表示当前请求需要相应的权限
		if re.MatchString(permissionKey) {
			if !utils.Contain(requiredPermission, allMethodPermission) {
				pass = false
			}
		}
	}
	// 如果路由没有与权限相关联，则默认允许访问
	if pass {
		context.Next()
		return
	}
	// 权限不足，返回错误
	utils.Response.NoAuth(context, "权限不足")
	context.Abort()
}

func WhiteList(context *gin.Context) bool {
	// 白名单
	whiteList := []string{config.Config.BASEURL + "/auth/login", config.Config.BASEURL + "/auth/captcha"}
	pass := false
	for _, v := range whiteList {
		if v == context.Request.URL.Path {
			context.Next()
			pass = true
			break
		}
	}
	return pass
}
