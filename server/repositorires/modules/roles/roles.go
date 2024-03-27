package rolesRepositorysModules

import (
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

func (r *rolesRepository) DeleteRole(id int) {}

func (r *rolesRepository) UpdateRole(role *responsies.UpdateRoleParams) {

}
func (r *rolesRepository) GetRoles(role *responsies.QueryRolesParams) {

}
