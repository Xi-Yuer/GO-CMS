package uploadTaskRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseUploadRoutes(r *gin.RouterGroup) {
	group := r.Group("/upload")
	{
		group.POST("/checkChunk", controllers.UploadController.CheckChunk)
		group.POST("/uploadChunk", controllers.UploadController.UploadChunk)
		group.POST("/mergeChunk", controllers.UploadController.MergeChunk)
	}
}
