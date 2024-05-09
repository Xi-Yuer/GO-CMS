package interfaceRpoterModuels

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseInterfaceRouter(r *gin.RouterGroup) {
	group := r.Group("/interface")
	{

		group.POST("", controllers.InterfaceController.CreateInterface)
		group.GET("", controllers.InterfaceController.GetAllInterface)
		group.DELETE("/:id", controllers.InterfaceController.DeleteInterface)
		group.PATCH("/:id", controllers.InterfaceController.UpdateInterface)
		group.GET("/page/:id", controllers.InterfaceController.GetInterfaceByPageID)
		group.GET("/role/:id", controllers.InterfaceController.GetInterfacesByRoleID)
	}
}
