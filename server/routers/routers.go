package routers

import (
	"github.com/Xi-Yuer/cms/config"
	"github.com/Xi-Yuer/cms/docs"
	"github.com/Xi-Yuer/cms/middlewares"
	authRouterModules "github.com/Xi-Yuer/cms/routers/modules/auth"
	usersRouterModules "github.com/Xi-Yuer/cms/routers/modules/users"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRouters() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	err := utils.Translator("zh")
	if err != nil {
		panic(err.Error())
	}

	v1 := r.Group("/api/v1", gin.Logger(), middlewares.AuthMiddleWareModule, middlewares.SessionMiddleWareModule(config.Config.APP.SESSIONSECRET))

	// swagger
	docs.SwaggerInfo.Title = "CMS API"
	docs.SwaggerInfo.Description = "This is a CMS server."
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	{
		usersRouterModules.UseUserRoutes(v1)
		authRouterModules.UseAuthRoutes(v1)
	}

	return r
}
