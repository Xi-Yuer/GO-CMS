package commitsRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseCommitsRoutes(r *gin.RouterGroup) {
	group := r.Group("/log")
	{
		group.GET("/commits", controllers.CommitsController.GetCommits)
		group.GET("/commits/count", controllers.CommitsController.GetCommitsCount)

	}
}
