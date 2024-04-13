package commitsControllerModules

import (
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

var CommitsController = &commitsController{}

type commitsController struct{}

// GetCommits 获取Git提交记录
// @Summary 获取Git提交记录
// @Description 获取Git提交记录
// @Tags 日志管理
// @Accept json
// @Produce json
// @Router /commits [get]
func (receiver *commitsController) GetCommits(context *gin.Context) {
	commits := repositories.CommitsRepositoryModules.GetCommits()
	utils.Response.Success(context, commits)
}
