package timeTaskRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseTimeTaskRoutes(r *gin.RouterGroup) {
	group := r.Group("/timeTask")

	{
		group.POST("/start/:id", controllers.TimeTaskController.StartTimeTask)
		group.POST("/stop/:id", controllers.TimeTaskController.StopTimeTask)
		group.PATCH("/update/:id", controllers.TimeTaskController.UpdateTimeTask)
		group.GET("", controllers.TimeTaskController.GetTimeTask)
	}
}
