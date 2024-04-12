package rolesAndPagesRepositoryModules

import (
	"database/sql"
	"fmt"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/utils"
)

var RolesAndPagesRepository = &rolesAndPagesRepository{}

type rolesAndPagesRepository struct {
}

func (r *rolesAndPagesRepository) CreateRecord(params *dto.CreateRolePermissionRecordParams) error {
	// 创建角色页面权限之前需要先将之前的记录全部删除，并且要保证事务的一致性
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	_, err = db.DB.Exec("DELETE FROM roles_pages WHERE role_id = ?", params.RoleID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	if len(params.PageID) == 0 {
		return nil
	}
	query := "INSERT INTO roles_pages (role_id, page_id) VALUES "
	for _, pageID := range params.PageID {
		query += fmt.Sprintf("('%s', '%s'),", params.RoleID, pageID)
	}
	query = query[:len(query)-1] // Remove the last comma and space

	_, err = db.DB.Exec(query)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *rolesAndPagesRepository) GetRecordsByRoleID(roleID string) ([]string, error) {
	rows, err := db.DB.Query("SELECT page_id FROM roles_pages WHERE role_id = ?", roleID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.Log.Error(err)
		}
	}(rows)

	var pageIDs []string

	for rows.Next() {
		var pageID string
		err := rows.Scan(&pageID)
		if err != nil {
			return nil, err
		}
		pageIDs = append(pageIDs, pageID)
	}

	return pageIDs, nil
}
