package authRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseAuthRoutes(r *gin.RouterGroup) {
	group := r.Group("/auth")
	{
		group.POST("/login", controllers.AuthController.Login)
		group.GET("/captcha", controllers.AuthController.Captcha)
		group.GET("/cookie", controllers.AuthController.Cookie)
	}
}
