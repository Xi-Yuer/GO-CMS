package rolesRepositorysModules

import (
	"database/sql"
	"errors"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
	"github.com/Xi-Yuer/cms/utils"
	"strings"
	"time"
)

var RolesRepository = &rolesRepository{}

type rolesRepository struct{}

func (r *rolesRepository) CreateRole(role *dto.CreateRoleParams) int64 {
	query := "INSERT INTO roles (role_id, role_name, description) VALUES (?, ?, ?)"
	roleID := utils.GenID()
	_, err := db.DB.Exec(query, roleID, role.RoleName, role.Description)
	if err != nil {
		return 0
	}
	return roleID
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
	_, err := db.DB.Exec(query, queryParams...)

	if err != nil {
		return err
	}

	return nil

}

func (r *rolesRepository) GetRoles(params *dto.QueryRolesParams) ([]*dto.SingleRoleResponse, error) {
	query := `
	SELECT roles.role_id, role_name, description,GROUP_CONCAT(roles_pages.page_id),GROUP_CONCAT(roles_interfaces.interface_id), create_time, update_time
	FROM roles
	LEFT JOIN roles_pages ON roles.role_id = roles_pages.role_id
	LEFT JOIN roles_interfaces ON roles.role_id = roles_interfaces.role_id
	WHERE delete_time IS NULL
	`

	var queryParams []interface{}

	if params.ID != "" {
		query += " AND roles.role_id = ?"
		queryParams = append(queryParams, params.ID)
	}
	if params.RoleName != "" {
		query += " AND role_name LIKE ?"
		queryParams = append(queryParams, "%"+params.RoleName+"%")
	}

	if params.Description != "" {
		query += " AND description LIKE ?"
		queryParams = append(queryParams, "%"+params.Description+"%")
	}

	if params.StartTime != "" {
		query += " AND create_time >= ?"
		queryParams = append(queryParams, params.StartTime)
	}

	if params.EndTime != "" {
		query += " AND create_time <= ?"
		queryParams = append(queryParams, params.EndTime)
	}

	query += " GROUP BY roles.role_id"
	query += ` LIMIT ?, ?`
	if params.Offset != 0 {
		queryParams = append(queryParams, params.Offset)
	} else {
		queryParams = append(queryParams, 0)
	}
	if params.Limit != 0 {
		queryParams = append(queryParams, params.Limit)
	} else {
		queryParams = append(queryParams, 10)
	}
	// 准备查询语句
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(queryParams...)
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
		var rolesID []uint8
		var interfaceID []uint8
		err := rows.Scan(&role.ID, &role.RoleName, &role.Description, &rolesID, &interfaceID, &role.CreateTime, &role.UpdateTime)
		if err != nil {
			return nil, err
		}
		if rolesID != nil {
			role.PagesID = strings.Split(string(rolesID), ",")
		}
		if interfaceID != nil {
			role.InterfacesID = strings.Split(string(interfaceID), ",")
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
		return errors.New("资源不存在")
	}

	return nil
}
