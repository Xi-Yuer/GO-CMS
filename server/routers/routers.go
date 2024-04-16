package routers

import (
	"github.com/Xi-Yuer/cms/config"
	"github.com/Xi-Yuer/cms/middlewares"
	authRouterModules "github.com/Xi-Yuer/cms/routers/modules/auth"
	commitsRouterModules "github.com/Xi-Yuer/cms/routers/modules/commits"
	departmentRouterModules "github.com/Xi-Yuer/cms/routers/modules/department"
	interfaceRpoterModuels "github.com/Xi-Yuer/cms/routers/modules/interface"
	logsRouterModules "github.com/Xi-Yuer/cms/routers/modules/logs"
	pagesRouterModules "github.com/Xi-Yuer/cms/routers/modules/pages"
	rolesRouterModules "github.com/Xi-Yuer/cms/routers/modules/roles"
	swaggerRouterModules "github.com/Xi-Yuer/cms/routers/modules/swagger"
	systemRouterModules "github.com/Xi-Yuer/cms/routers/modules/system"
	timeTaskRouterModules "github.com/Xi-Yuer/cms/routers/modules/timeTask"
	uploadTaskRouterModules "github.com/Xi-Yuer/cms/routers/modules/upload"
	usersRouterModules "github.com/Xi-Yuer/cms/routers/modules/users"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

func SetUpRouters() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	err := utils.Translator("zh")
	if err != nil {
		panic(err.Error())
	}

	v1 := r.Group(
		config.Config.APP.BASEURL,
		gin.Logger(),
		middlewares.LogsMiddlewareModule.SystemLogMiddleware,
		middlewares.AuthMiddleWareModule,
		middlewares.AuthMethodMiddleWare,
		middlewares.SessionMiddleWareModule(config.Config.APP.SESSIONSECRET),
	)

	{
		usersRouterModules.UseUserRoutes(v1)
		authRouterModules.UseAuthRoutes(v1)
		rolesRouterModules.UseRolesRoutes(v1)
		pagesRouterModules.UsePagesRoutes(v1)
		departmentRouterModules.UseDepartmentRoutes(v1)
		interfaceRpoterModuels.UseInterfaceRouter(v1)
		logsRouterModules.UseLogRoutes(v1)
		commitsRouterModules.UseCommitsRoutes(v1)
		systemRouterModules.UseSystemRoutes(v1)
		timeTaskRouterModules.UseTimeTaskRoutes(v1)
		uploadTaskRouterModules.UseUploadRoutes(v1)
		swaggerRouterModules.UseSwaggerRoutes(r)
	}

	return r
}
