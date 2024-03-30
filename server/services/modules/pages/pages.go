package pagesServiceModules

import (
	"github.com/Xi-Yuer/cms/dto"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
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
