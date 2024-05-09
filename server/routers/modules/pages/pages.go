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
		group.GET("", controllers.PagesController.GetPages)
		group.GET("/user", controllers.PagesController.GetUserMenus)
		group.GET("/role/:id", controllers.PagesController.GetPagesByRoleID)
		group.PATCH("/:id", controllers.PagesController.UpdatePage)
	}
}
