package rolesRepositorysModules

import (
	"database/sql"
	"fmt"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/responsies"
	"github.com/Xi-Yuer/cms/utils"
)

var RolesRepository = &rolesRepository{}

type rolesRepository struct{}

func (r *rolesRepository) CreateRole(role *responsies.CreateRoleParams) error {
	query := "INSERT INTO roles (role_id, role_name, description) VALUES (?, ?, ?)"
	roleID := utils.GenID()
	_, err := db.DB.Exec(query, roleID, role.RoleName, role.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *rolesRepository) DeleteRole(id string) error {
	query := "DELETE FROM roles WHERE role_id = ?"
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *rolesRepository) UpdateRole(role *responsies.UpdateRoleParams) error {
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
	queryParams = append(queryParams, role.ID)
	fmt.Println(query)
	fmt.Println(queryParams)

	_, err := db.DB.Exec(query, queryParams...)

	if err != nil {
		return err
	}

	return nil

}
func (r *rolesRepository) GetRoles() ([]*responsies.SingleRoleResponse, error) {
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

	roles := make([]*responsies.SingleRoleResponse, 0)
	for rows.Next() {
		role := &responsies.SingleRoleResponse{}
		err := rows.Scan(&role.ID, &role.RoleName, &role.Description, &role.CreateTime, &role.UpdateTime)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (r *rolesRepository) FindRoleById(id string) *responsies.SingleRoleResponse {
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
	role := &responsies.SingleRoleResponse{}
	for rows.Next() {
		err := rows.Scan(&role.ID, &role.RoleName, &role.Description, &role.CreateTime, &role.UpdateTime)
		if err != nil {
			utils.Log.Error(err)
		}
	}
	return role
}
