package rolesRepositorysModules

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/utils"
	"strings"
	"time"
)

var RolesRepository = &rolesRepository{}

type rolesRepository struct{}

func (r *rolesRepository) CreateRole(role *dto.CreateRoleParams) error {
	query := "INSERT INTO roles (role_id, role_name, description) VALUES (?, ?, ?)"
	roleID := utils.GenID()
	_, err := db.DB.Exec(query, roleID, role.RoleName, role.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *rolesRepository) DeleteRole(id string) error {
	query := "UPDATE roles SET delete_time = ? WHERE role_id = ?"
	_, err := db.DB.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (r *rolesRepository) UpdateRole(role *dto.UpdateRoleParams, id string) error {
	query := "UPDATE roles SET "
	var (
		queryParams []any
		hasSet      bool
	)
	if role.RoleName != "" {
		queryParams = append(queryParams, role.RoleName)
		query += "role_name = ?, "
		hasSet = true
	}

	if role.Description != "" {
		queryParams = append(queryParams, role.Description)
		query += "description = ?, "
		hasSet = true
	}
	if !hasSet {
		return nil
	}

	// 去除最后一个逗号和空格
	query = query[:len(query)-2]

	query += " WHERE role_id = ?"
	queryParams = append(queryParams, id)
	fmt.Println(query)
	fmt.Println(queryParams)

	_, err := db.DB.Exec(query, queryParams...)

	if err != nil {
		return err
	}

	return nil

}

func (r *rolesRepository) GetRoles() ([]*dto.SingleRoleResponse, error) {
	query := "SELECT role_id, role_name, description, create_time, update_time FROM roles WHERE delete_time IS NULL "

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	roles := make([]*dto.SingleRoleResponse, 0)
	for rows.Next() {
		role := &dto.SingleRoleResponse{}
		err := rows.Scan(&role.ID, &role.RoleName, &role.Description, &role.CreateTime, &role.UpdateTime)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (r *rolesRepository) FindRoleById(id string) *dto.SingleRoleResponse {
	query := "SELECT role_id, role_name, description, create_time, update_time FROM roles WHERE role_id = ? AND delete_time IS NULL"
	rows, err := db.DB.Query(query, id)
	if err != nil {
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			utils.Log.Error(err)
		}
	}(rows)
	role := &dto.SingleRoleResponse{}
	for rows.Next() {
		err := rows.Scan(&role.ID, &role.RoleName, &role.Description, &role.CreateTime, &role.UpdateTime)
		if err != nil {
			utils.Log.Error(err)
		}
	}
	if role.ID == "" {
		return nil
	}
	return role
}

// CheckRolesExistence 检查角色是否都存在
func (r *rolesRepository) CheckRolesExistence(roleIDs []string) error {
	// 构建 IN 子句
	var placeholders []string
	var args []interface{}
	for _, id := range roleIDs {
		placeholders = append(placeholders, "?")
		args = append(args, id)
	}
	query := "SELECT COUNT(*) FROM roles WHERE delete_time IS NULL AND role_id IN (" + strings.Join(placeholders, ",") + ") "

	row := db.DB.QueryRow(query, args...)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return err
	}
	// 判断是否所有角色都存在
	if count != len(roleIDs) {
		return errors.New("角色不存在")
	}

	return nil
}
