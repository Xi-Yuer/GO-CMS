package InterfaceRpoterModuels

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseInterfaceRouter(r *gin.RouterGroup) {
	group := r.Group("/interface")
	{

		group.POST("/", controllers.InterfaceController.CreateInterface)
		group.DELETE("/:id", controllers.InterfaceController.DeleteInterface)
		group.PATCH("/:id", controllers.InterfaceController.UpdateInterface)
		group.GET("/:id", controllers.InterfaceController.GetInterfaceByPageID)
	}
}
