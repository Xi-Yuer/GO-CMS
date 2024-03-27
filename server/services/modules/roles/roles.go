package rolesServiceModules

import (
	repositories "github.com/Xi-Yuer/cms/repositorires/modules"
	"github.com/Xi-Yuer/cms/responsies"
)

var RolesService = &rolesService{}

type rolesService struct{}

func (r *rolesService) CreateRole(role *responsies.CreateRoleParams) error {
	return repositories.RoleRepositorysModules.CreateRole(role)
}

func (r *rolesService) DeleteRole(id int) {}

func (r *rolesService) UpdateRole(role *responsies.UpdateRoleParams) {

}
func (r *rolesService) GetRoles(role *responsies.QueryRolesParams) {

}
