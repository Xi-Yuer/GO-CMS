package pagesRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UsePagesRoutes(r *gin.RouterGroup) {
	group := r.Group("/pages")
	{
		group.POST("", controllers.PagesController.CreatePage)
		group.DELETE("/:id", controllers.PagesController.DeletePage)
	}
}
