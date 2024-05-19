package templateControllerModules

import (
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/services"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

var TemplateController = &templateController{}

type templateController struct{}

func (t *templateController) CreateTemplate(context *gin.Context) {
	var params dto.CreateTemplateRequestParams
	if err := context.ShouldBind(&params); err != nil {
		utils.Response.ParameterTypeError(context, err.Error())
		return
	}
	parseTemplate, err := services.TemplateService.CreateTemplate(&params)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, parseTemplate)
}
