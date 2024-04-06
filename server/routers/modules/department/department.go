package departmentRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseDepartmentRoutes(r *gin.RouterGroup) {
	group := r.Group("/department")
	group.POST("", controllers.DepartmentController.CreateDepartment)
	group.DELETE("/:id", controllers.DepartmentController.DeleteDepartment)
	group.GET("", controllers.DepartmentController.GetDepartments)
	group.PATCH("/:id", controllers.DepartmentController.UpdateDepartment)
}
