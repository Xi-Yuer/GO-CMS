package usersRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseUserRoutes(r *gin.RouterGroup) {

	userRouters := r.Group("/users")

	{
		userRouters.POST("", controllers.UserController.CreateUser)
		userRouters.GET("/:id", controllers.UserController.GetUser)
		userRouters.GET("", controllers.UserController.GetUsers)
		userRouters.GET("/search", controllers.UserController.FindUserByParams)
		userRouters.PATCH("/:id", controllers.UserController.UpdateUser)
		userRouters.DELETE("/:id", controllers.UserController.DeleteUser)
	}
}
