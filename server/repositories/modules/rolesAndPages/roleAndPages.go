package rolesAndPagesRepositoryModules

import (
	"fmt"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
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
