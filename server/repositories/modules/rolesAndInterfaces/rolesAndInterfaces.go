package rolesAndInterfacesRepositoryModules

import (
	"fmt"
	"github.com/Xi-Yuer/cms/db"
	"github.com/Xi-Yuer/cms/dto"
)

var RolesAndInterfacesRepository = &rolesAndInterfacesRepository{}

type rolesAndInterfacesRepository struct{}

func (r *rolesAndInterfacesRepository) CreateRecord(params *dto.CreateRolePermissionRecordParams) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	// 给角色分配接口权限，需要清空之间的数据
	query := `DELETE FROM roles_interfaces WHERE role_id = ?`

	if _, err := db.DB.Exec(query, params.RoleID); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	// 接口ID为空
	if len(params.InterfaceID) == 0 {
		return nil
	}
	query = `INSERT INTO roles_interfaces (role_id, interface_id) VALUES `
	for _, inter := range params.InterfaceID {
		query += fmt.Sprintf("('%s', '%s'),", params.RoleID, inter)
	}
	query = query[:len(query)-1] // Remove the last comma and space

	if _, err := db.DB.Exec(query); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	return nil
}
func (r *rolesAndInterfacesRepository) GetRecordByRoleID(roleID string) ([]string, error) {
	query := "SELECT interface_id FROM roles_interfaces WHERE role_id = ?"
	rows, err := db.DB.Query(query, roleID)
	if err != nil {
		return nil, err
	}
	var interfaceIDs []string
	for rows.Next() {
		var interfaceID string
		if err := rows.Scan(&interfaceID); err != nil {
			return nil, err
		}
		interfaceIDs = append(interfaceIDs, interfaceID)

	}
	return interfaceIDs, nil
}
