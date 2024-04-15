package timeTaskControllerModules

import (
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/services"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

var TimeTaskController = &timeTaskController{}

type timeTaskController struct{}

// GetTimeTask 获取定时任务
// @Summary 获取定时任务
// @Description 获取定时任务
// @Tags 定时任务
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /timeTask [get]
func (t *timeTaskController) GetTimeTask(context *gin.Context) {
	tasks := services.TimeTaskService.GetTimeTask()
	utils.Response.Success(context, tasks)
}

// StartTimeTask 开启定时任务
// @Summary 开启定时任务
// @Description 开启定时任务
// @Tags 定时任务
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /timeTask/start/{id} [post]
func (t *timeTaskController) StartTimeTask(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterMissing(context, "id不能为空")
		return
	}
	if err := services.TimeTaskService.StartTimeTask(id); err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, "启动定时任务成功")
}

// StopTimeTask 停止定时任务
// @Summary 停止定时任务
// @Description 停止定时任务
// @Tags 定时任务
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /timeTask/stop/{id} [post]
func (t *timeTaskController) StopTimeTask(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterMissing(context, "id不能为空")
		return
	}
	if err := services.TimeTaskService.StopTimeTask(id); err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, "停止定时任务成功")
}

// UpdateTimeTask 更新定时任务
// @Summary 更新定时任务
// @Description 更新定时任务
// @Tags 定时任务
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /timeTask/update/{id} [patch]
func (t *timeTaskController) UpdateTimeTask(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		utils.Response.ParameterMissing(context, "id不能为空")
		return
	}

	var params dto.UpdateTimeTask
	if err := context.ShouldBind(&params); err != nil {
		utils.Response.ParameterMissing(context, err.Error())
		return
	}

	if err := services.TimeTaskService.UpdateTask(id, &params); err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, "更新定时任务成功")
}
