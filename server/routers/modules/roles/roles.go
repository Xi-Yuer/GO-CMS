package rolesRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseRolesRoutes(r *gin.RouterGroup) {
	group := r.Group("/roles")
	{
		group.POST("/query", controllers.RoleController.GetRoles)
		group.POST("", controllers.RoleController.CreateRole)
		group.DELETE("/:id", controllers.RoleController.DeleteRole)
		group.PATCH("/:id", controllers.RoleController.UpdateRole)
		group.POST("/bindUser", controllers.RoleController.CreateOneRecord)
		group.POST("/deBindUser", controllers.RoleController.DeleteOneRecord)
		group.POST("/export", controllers.RoleController.ExportExcel)
	}
}
