package pagesRepositoryModules

import (
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/utils"
)

var PageRepository = &pageRepository{}

type pageRepository struct {
}

func (p *pageRepository) CreatePage(params *dto.CreatePageParams) error {
	var err error
	if params.ParentPage == "" {
		query := `INSERT INTO pages (page_id, page_name, page_path, page_icon, page_component, can_edit) VALUES (?,?,?,?,?,?)`
		pageID := utils.GenID()
		_, err = db.DB.Exec(query, pageID, params.PageName, params.PagePath, params.PageIcon, params.PageComponent, 1)
	} else {
		query := `INSERT INTO pages (page_id, page_name, page_path, page_icon, page_component, parent_page, can_edit) VALUES (?,?,?,?,?,?,?)`
		pageID := utils.GenID()
		_, err = db.DB.Exec(query, pageID, params.PageName, params.PagePath, params.PageIcon, params.PageComponent, params.ParentPage, 1)
	}
	if err != nil {
		return err
	}
	return nil
}

func (p *pageRepository) DeletePage(pageID string) error {
	query := `DELETE FROM pages WHERE page_id = ?`
	_, err := db.DB.Exec(query, pageID)
	return err
}
