package systemRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseSystemRoutes(r *gin.RouterGroup) {
	group := r.Group("/system")
	{
		group.GET("", controllers.SystemController.GetSystemInfo)
	}
}
