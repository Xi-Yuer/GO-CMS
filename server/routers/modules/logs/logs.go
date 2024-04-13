package logsRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseLogRoutes(r *gin.RouterGroup) {
	group := r.Group("/log")
	{
		group.GET("/system", controllers.LogsController.GetLogRecords)
	}
}
