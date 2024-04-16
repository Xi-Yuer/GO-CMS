package uploadControllerModules

import (
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/services"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

var UploadController = &uploadController{}

type uploadController struct{}

func (u *uploadController) CheckChunk(context *gin.Context) {
	var params dto.UploadBigFileRequest
	if err := context.ShouldBind(&params); err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	chunk, err := services.UploadService.CheckChunk(&params)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	if len(chunk.ChunkList) == 0 {
		context.JSON(201, gin.H{
			"code": 201,
			"msg":  "没有分片",
			"data": nil,
		})
		return
	}
	utils.Response.Success(context, chunk)
}
func (u *uploadController) UploadChunk(context *gin.Context) {
	var params dto.UploadBigFileRequest
	if err := context.ShouldBind(&params); err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}

	if err := services.UploadService.UploadChunk(context, &params); err != nil {
		return
	}
	utils.Response.Success(context, nil)
}
func (u *uploadController) MergeChunk(context *gin.Context) {
	var params dto.UploadBigFileRequest
	if err := context.ShouldBind(&params); err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	chunk, err := services.UploadService.MergeChunk(params)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, chunk)
}
