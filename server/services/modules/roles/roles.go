package rolesServiceModules

import (
	"errors"
	"github.com/Xi-Yuer/cms/dto"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
)

var RolesService = &rolesService{}

type rolesService struct{}

func (r *rolesService) CreateRole(role *dto.CreateRoleParams) error {
	return repositories.RoleRepositorysModules.CreateRole(role)
}

func (r *rolesService) DeleteRole(id string) error {
	singleRoleResponse := repositories.RoleRepositorysModules.FindRoleById(id)
	if singleRoleResponse.ID == "" {
		return errors.New("角色不存在")
	}
	return repositories.RoleRepositorysModules.DeleteRole(id)
}

func (r *rolesService) UpdateRole(role *dto.UpdateRoleParams, id string) error {
	singleRoleResponse := repositories.RoleRepositorysModules.FindRoleById(id)
	if singleRoleResponse.ID == "" {
		return errors.New("角色不存在")
	}
	return repositories.RoleRepositorysModules.UpdateRole(role, id)

}
func (r *rolesService) GetRoles() ([]*dto.SingleRoleResponse, error) {
	return repositories.RoleRepositorysModules.GetRoles()
}
