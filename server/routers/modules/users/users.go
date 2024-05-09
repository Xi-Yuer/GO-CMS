package usersRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseUserRoutes(r *gin.RouterGroup) {

	group := r.Group("/users")

	{
		group.POST("", controllers.UserController.CreateUser)
		group.GET("/:id", controllers.UserController.GetUser)
		group.POST("/query", controllers.UserController.GetUsers)
		group.POST("/query/role/:id", controllers.UserController.FindUserByParamsAndOutRoleID)
		group.GET("/role/:id", controllers.UserController.GetUserByRoleID)
		group.PATCH("/:id", controllers.UserController.UpdateUser)
		group.DELETE("/:id", controllers.UserController.DeleteUser)
		group.POST("/export", controllers.UserController.ExportExcel)
	}
}
