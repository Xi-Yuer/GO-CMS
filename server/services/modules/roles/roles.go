package rolesServiceModules

import (
	"errors"
	repositories "github.com/Xi-Yuer/cms/repositorires/modules"
	"github.com/Xi-Yuer/cms/responsies"
)

var RolesService = &rolesService{}

type rolesService struct{}

func (r *rolesService) CreateRole(role *responsies.CreateRoleParams) error {
	return repositories.RoleRepositorysModules.CreateRole(role)
}

func (r *rolesService) DeleteRole(id string) error {
	singleRoleResponse := repositories.RoleRepositorysModules.FindRoleById(id)
	if singleRoleResponse.ID == "" {
		return errors.New("角色不存在")
	}
	return repositories.RoleRepositorysModules.DeleteRole(id)
}

func (r *rolesService) UpdateRole(role *responsies.UpdateRoleParams) error {
	singleRoleResponse := repositories.RoleRepositorysModules.FindRoleById(role.ID)
	if singleRoleResponse.ID == "" {
		return errors.New("角色不存在")
	}
	return repositories.RoleRepositorysModules.UpdateRole(role)

}
func (r *rolesService) GetRoles() ([]*responsies.SingleRoleResponse, error) {
	return repositories.RoleRepositorysModules.GetRoles()
}
