package rolesRepositorysModules

import "github.com/Xi-Yuer/cms/responsies"

var RolesRepository = &rolesRepository{}

type rolesRepository struct{}

func (r *rolesRepository) CreateRole(role *responsies.CreateRoleParams) error {
	return nil
}

func (r *rolesRepository) DeleteRole(id int) {}

func (r *rolesRepository) UpdateRole(role *responsies.UpdateRoleParams) {

}
func (r *rolesRepository) GetRoles(role *responsies.QueryRolesParams) {

}
