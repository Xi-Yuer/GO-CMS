package routers

import (
	"github.com/Xi-Yuer/cms/docs"
	"github.com/Xi-Yuer/cms/middlewares"
	routers "github.com/Xi-Yuer/cms/routers/modules"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRouters() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	v1 := r.Group("/api/v1", middlewares.RequestMiddlewareModule.RequestLoggerMiddleware)

	// swagger
	docs.SwaggerInfo.Title = "CMS API"
	docs.SwaggerInfo.Description = "This is a CMS server."
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	{
		routers.UseUserRoutes(v1)
	}

	return r
}
