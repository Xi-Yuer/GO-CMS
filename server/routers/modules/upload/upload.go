package uploadTaskRouterModules

import (
	"github.com/Xi-Yuer/cms/controllers"
	"github.com/Xi-Yuer/cms/middlewares"
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
		// 通过 GET 下载文件需要验证 Cookie 信息，请求请求文件下载的时候请求一下 Cookie 信息
		group.GET("/download/aHref/:id", middlewares.AuthVerifyCookie, controllers.UploadController.DownloadFile)
	}
}
