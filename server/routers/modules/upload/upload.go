package uploadTaskRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/gin-gonic/gin"
)

func UseUploadRoutes(r *gin.RouterGroup) {
	group := r.Group("/upload")
	{
		group.POST("/check", controllers.UploadController.CheckFile)
		group.POST("", controllers.UploadController.UploadChunk)
		group.POST("/finish", controllers.UploadController.FinishUpload)
		group.DELETE("/del/:id", controllers.UploadController.DeleteFile)
		group.GET("", controllers.UploadController.GetFile)
		group.POST("/download/:id", controllers.UploadController.DownloadFile)
	}
}
