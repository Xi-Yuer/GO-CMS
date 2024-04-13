package commitsRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseCommitsRoutes(r *gin.RouterGroup) {
	group := r.Group("/commits")
	{
		group.GET("", controllers.CommitsController.GetCommits)
	}
}
