package pagesServiceModules

import (
	"github.com/Xi-Yuer/cms/dto"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
	"github.com/Xi-Yuer/cms/utils"
)

var PageService = &pageService{}

type pageService struct {
}

func (p *pageService) CreatePage(params *dto.CreatePageParams) error {
	return repositories.PageRepositorysModules.CreatePage(params)
}

func (p *pageService) DeletePage(params string) error {
	return repositories.PageRepositorysModules.DeletePage(params)
}

func (p *pageService) GetPages() ([]*dto.SinglePageResponse, error) {
	pages, err := repositories.PageRepositorysModules.GetPages()
	if err != nil {
		return nil, err
	}
	buildPages := utils.BuildPages(pages)
	return buildPages, nil
}
