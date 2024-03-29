package rolesRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseRolesRoutes(r *gin.RouterGroup) {
	group := r.Group("/roles")
	{
		group.POST("/", controllers.RoleController.CreateRole)
		group.DELETE("/:id", controllers.RoleController.DeleteRole)
		group.PUT("/:id", controllers.RoleController.UpdateRole)
		group.GET("/", controllers.RoleController.GetRoles)
	}
}
