package logsControllerModules

import (
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/services"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

var LogsController = &logsControllerModules{}

type logsControllerModules struct{}

// GetLogRecords 获取日志
// @Summary 获取日志
// @Description 获取日志
// @Tags 日志管理
// @Accept json
// @Produce json
// @Router /logs [get]
func (l *logsControllerModules) GetLogRecords(context *gin.Context) {
	var params dto.Page
	if err := context.ShouldBind(&params); err != nil {
		utils.Response.ParameterMissing(context, err.Error())
		return
	}
	logRecords, err := services.LogsService.GetLogRecords(&params)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, logRecords)
}
