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
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /log/commits [get]
func (c *commitsController) GetCommits(context *gin.Context) {
	commits := repositories.CommitsRepositoryModules.GetCommits()
	// 分类提交记录到日期
	groupedCommits := utils.GroupCommitsByDate(commits)

	// 格式化为所需的结构
	formattedCommits := utils.FormatCommits(groupedCommits)
	utils.Response.Success(context, formattedCommits)
}

// GetCommitsCount 获取Git提交次数
// @Summary 获取Git提交次数
// @Description 获取Git提交次数
// @Tags 日志管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /log/commits/count [get]
func (c *commitsController) GetCommitsCount(context *gin.Context) {
	count, err := repositories.CommitsRepositoryModules.GetCommitsCount()
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, count)
}
