package pagesRepositoryModules

import (
	"database/sql"
	"errors"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/utils"
	"strings"
)

var PageRepository = &pageRepository{}

type pageRepository struct {
}

func (p *pageRepository) CreatePage(params *dto.CreatePageParams) error {
	var err error
	pageID := utils.GenID()
	if params.ParentPage == "" {
		query := `INSERT INTO pages (page_id, page_name, page_order, page_path, page_icon, page_component, can_edit) VALUES (?,?,?,?,?,?,?)`
		_, err = db.DB.Exec(query, pageID, params.PageName, params.PageOrder, params.PagePath, params.PageIcon, params.PageComponent, 1)
	} else {
		query := `INSERT INTO pages (page_id, page_name,page_order, page_path, page_icon, page_component, parent_page, can_edit) VALUES (?,?,?,?,?,?,?,?)`
		_, err = db.DB.Exec(query, pageID, params.PageName, params.PageOrder, params.PagePath, params.PageIcon, params.PageComponent, params.ParentPage, 1)
	}
	if err != nil {
		return err
	}
	return nil
}

func (p *pageRepository) DeletePage(pageID string) error {
	page, err := p.FindPageByID(pageID)
	if err != nil {
		return err
	}
	if page.CanEdit == 0 {
		return errors.New("该页面不可删除")
	}
	query := `DELETE FROM pages WHERE page_id = ?`
	_, err = db.DB.Exec(query, pageID)
	return err
}

func (p *pageRepository) FindPageByID(id string) (*dto.SinglePageResponse, error) {
	query := "SELECT page_id, page_name, page_order, page_path, page_icon, page_component, parent_page, can_edit, create_time, update_time FROM pages WHERE page_id = ?"
	var Page dto.SinglePageResponse
	err := db.DB.QueryRow(query, id).Scan(&Page.PageID, &Page.PageName, &Page.PageOrder, &Page.PagePath, &Page.PageIcon, &Page.PageComponent, &Page.ParentPage, &Page.CanEdit, &Page.CreatedTime, &Page.UpdateTime)
	if err != nil {
		return nil, err
	}
	return &Page, nil
}

func (p *pageRepository) FindChildPagesByParentID(parentID string) ([]dto.SinglePageResponse, error) {
	query := "SELECT page_id, page_name, page_order, page_path, page_icon, page_component, parent_page, can_edit, create_time, update_time FROM pages WHERE parent_page = ?"
	var Page []dto.SinglePageResponse
	exec, err := db.DB.Query(query, parentID)
	if err != nil {
		return nil, err
	}
	for exec.Next() {
		var page dto.SinglePageResponse
		err := exec.Scan(&page.PageID, &page.PageName, &page.PageOrder, &page.PagePath, &page.PageIcon, &page.PageComponent, &page.ParentPage, &page.CanEdit, &page.CreatedTime, &page.UpdateTime)
		if err != nil {
			return nil, err
		}
		Page = append(Page, page)
	}
	return Page, nil
}

func (p *pageRepository) CheckPagesExistence(pagesIDs []string) error {
	// 构建 IN 子句
	var placeholders []string
	var args []interface{}
	for _, id := range pagesIDs {
		placeholders = append(placeholders, "?")
		args = append(args, id)
	}
	query := "SELECT COUNT(*) FROM pages WHERE delete_time IS NULL AND pages.page_id IN (" + strings.Join(placeholders, ",") + ") "

	row := db.DB.QueryRow(query, args...)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return err
	}
	// 判断是否所有角色都存在
	if count != len(pagesIDs) {
		return errors.New("页面不存在")
	}

	return nil
}

func (p *pageRepository) GetPages() ([]*dto.SinglePageResponse, error) {
	query := "SELECT page_id, page_name, page_order, page_path, page_icon, page_component, parent_page, can_edit, create_time, update_time FROM pages WHERE delete_time IS NULL"

	exec, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer func(exec *sql.Rows) {
		err := exec.Close()
		if err != nil {
			utils.Log.Error(err)
		}
	}(exec)

	var Pages []*dto.SinglePageResponse
	for exec.Next() {
		page := dto.SinglePageResponse{}
		err := exec.Scan(&page.PageID, &page.PageName, &page.PageOrder, &page.PagePath, &page.PageIcon, &page.PageComponent, &page.ParentPage, &page.CanEdit, &page.CreatedTime, &page.UpdateTime)
		if err != nil {
			return nil, err
		}
		Pages = append(Pages, &page)
	}

	return Pages, nil
}
