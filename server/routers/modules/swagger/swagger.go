package swaggerRouterModules

import (
	"github.com/Xi-Yuer/cms/config"
	"github.com/Xi-Yuer/cms/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func UseSwaggerRoutes(r *gin.Engine) {
	docs.SwaggerInfo.Title = "后台管理系统API接口"
	docs.SwaggerInfo.Description = "后台管理系统API接口"
	docs.SwaggerInfo.BasePath = config.Config.BASEURL
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
