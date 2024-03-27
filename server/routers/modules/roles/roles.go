package rolesRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseRolesRoutes(r *gin.RouterGroup) {
	group := r.Group("/role")
	{
		group.POST("/", controllers.RoleController.CreateRole)
	}
}
