package pagesControllerModules

import (
	"github.com/Xi-Yuer/cms/constant"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/services"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

var PagesController = &pagesController{}

type pagesController struct {
}

// CreatePage 创建菜单
// @Summary 创建菜单
// @Description 创建菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Router /pages [post]
func (p *pagesController) CreatePage(context *gin.Context) {
	var pageParams dto.CreatePageParams
	if err := context.ShouldBind(&pageParams); err != nil {
		utils.Response.ParameterMissing(context, err.Error())
	}
	err := services.PageService.CreatePage(&pageParams)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, nil)
}

// DeletePage 删除菜单
// @Summary 删除菜单
// @Description 删除菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Router /pages/{id} [delete]
func (p *pagesController) DeletePage(context *gin.Context) {
	param := context.Param("id")
	if param == "" {
		utils.Response.ParameterMissing(context, "pageID不能为空")
		return
	}
	err := services.PageService.DeletePage(param)
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, nil)
}

// GetPages 获取菜单
// @Summary 获取菜单
// @Description 获取菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Router /pages [get]
func (p *pagesController) GetPages(context *gin.Context) {
	pages, err := services.PageService.GetPages()
	if err != nil {
		utils.Response.ServerError(context, err.Error())
		return
	}
	utils.Response.Success(context, pages)
}

// GetUserMenus 获取用户菜单
// @Summary 获取用户菜单
// @Description 获取用户菜单
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /pages/menu [get]
func (a *pagesController) GetUserMenus(context *gin.Context) {
	jwtPayload, exist := context.Get(constant.JWTPAYLOAD)
	if !exist {
		utils.Response.ParameterMissing(context, "用户id不能为空")
		return
	}
	menus, err := services.PageService.GetUserMenus(jwtPayload.(*dto.JWTPayload).ID)
	if err != nil {
		return
	}
	utils.Response.Success(context, menus)
}
