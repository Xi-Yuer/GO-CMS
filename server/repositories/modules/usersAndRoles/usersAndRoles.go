package usersAndRolesRepositorysModules

import (
	"database/sql"
	"fmt"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/utils"
)

var UsersAndRolesRepositorys = &usersAndRolesRepositorys{}

type usersAndRolesRepositorys struct {
}

func (u *usersAndRolesRepositorys) CreateRecords(userID string, roleID []string) error {
	// 创建用户角色之前需要先将之前的记录全部删除，并且要保证事务的一致性
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	_, err = db.DB.Exec("DELETE FROM users_roles WHERE user_id = ?", userID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	query := "INSERT INTO users_roles (user_id, role_id) VALUES "
	for _, roleID := range roleID {
		query += fmt.Sprintf("('%s', '%s'),", userID, roleID)
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

func (u *usersAndRolesRepositorys) FindUserRolesHasDetail(userID string) ([]*dto.SingleRoleResponse, error) {
	// 连表查询
	query := "SELECT r.role_id, role_name,description,create_time,update_time FROM users_roles JOIN cms.roles r on r.role_id = users_roles.role_id WHERE user_id = ? AND delete_time IS NULL"
	rows, err := db.DB.Query(query, userID)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.Log.Error(err)
		}
	}(rows)

	var roles = make([]*dto.SingleRoleResponse, 0)
	for rows.Next() {
		role := &dto.SingleRoleResponse{}
		err := rows.Scan(&role.ID, &role.RoleName, &role.Description, &role.CreateTime, &role.UpdateTime)
		if err != nil {
			continue
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (u *usersAndRolesRepositorys) FindUserRolesID(userID string) []string {
	query := "SELECT role_id FROM users_roles WHERE user_id = ?"
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		utils.Log.Error(err)
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.Log.Error(err)
		}
	}(rows)

	var roleIDs []string
	for rows.Next() {
		var roleID string
		err := rows.Scan(&roleID)
		if err != nil {
			continue
		}
		roleIDs = append(roleIDs, roleID)
	}
	return roleIDs
}

func (u *usersAndRolesRepositorys) CreateOneRecord(params *dto.CreateOneRecord) error {
	query := "INSERT INTO users_roles (user_id, role_id) VALUES (?, ?)"
	if _, err := db.DB.Exec(query, params.UserID, params.RoleID); err != nil {
		return err
	}
	return nil
}

func (u *usersAndRolesRepositorys) DeleteOneRecord(params *dto.DeleteOneRecord) error {
	query := "DELETE FROM users_roles WHERE user_id = ? AND role_id = ?"
	if _, err := db.DB.Exec(query, params.UserID, params.RoleID); err != nil {
		return err
	}
	return nil
}
