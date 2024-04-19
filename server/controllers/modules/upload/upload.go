package uploadControllerModules

import (
	"github.com/Xi-Yuer/cms/config"
	"github.com/Xi-Yuer/cms/constant"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/services"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

var UploadController = &uploadController{}

type uploadController struct{}

// CheckFile 检查文件状态（是否上传过、之前上传了多少）
// @Summary 检查文件状态
// @Description 检查文件状态
// @Tags 文件管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /upload/check [post]
func (u *uploadController) CheckFile(context *gin.Context) {
	var params dto.CheckChunkRequest
	if err := context.ShouldBind(&params); err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	chunk, err := services.UploadService.CheckFile(&params)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, chunk)
}

// UploadChunk 上传文件
// @Summary 上传文件
// @Description 上传文件
// @Tags 文件管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /upload [post]
func (u *uploadController) UploadChunk(context *gin.Context) {
	var params dto.UploadBigFileRequest
	if err := context.ShouldBind(&params); err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}

	if err := services.UploadService.UploadChunk(&params); err != nil {
		return
	}
	utils.Response.Success(context, nil)
}

// FinishUpload 完成上传
// @Summary 完成上传
// @Description 完成上传
// @Tags 文件管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /upload/finish [post]
func (u *uploadController) FinishUpload(context *gin.Context) {
	var params dto.UploadFinishRequest
	if err := context.ShouldBind(&params); err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	value, _ := context.Get(constant.JWTPAYLOAD)
	userAccount := value.(*dto.JWTPayload)
	chunk, err := services.UploadService.FinishUpload(userAccount.Account, &params)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, chunk)
}

// DeleteFile 删除文件
// @Summary 删除文件
// @Description 删除文件
// @Tags 文件管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /upload/del/{id} [delete]
func (u *uploadController) DeleteFile(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterTypeError(context, "id不能为空")
		return
	}
	if err := services.UploadService.DeleteFile(id); err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, nil)
}

// GetFile 获取文件
// @Summary 获取文件
// @Description 获取文件
// @Tags 文件管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /upload [get]
func (u *uploadController) GetFile(context *gin.Context) {
	var params dto.Page
	if err := context.ShouldBind(&params); err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}

	fileList, err := services.UploadService.GetFileList(&params)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, fileList)
}

// DownloadFile 文件下载
// @Summary 文件下载
// @Description 文件下载
// @Tags 文件管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /upload/aHref/download [post]
func (u *uploadController) DownloadFile(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterTypeError(context, "id不能为空")
		return
	}
	file, err := services.UploadService.DownloadFile(id)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	context.Header("Content-Type", "application/octet-stream")
	context.Header("Content-Disposition", "attachment; filename="+file.FileName)
	context.Header("Content-Transfer-Encoding", "binary")
	context.File(config.Config.FILEPATH + id)
}
