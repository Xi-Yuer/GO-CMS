package rolesServiceModules

import (
	"errors"
	"github.com/Xi-Yuer/cms/dto"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
	"strconv"
)

var RolesService = &rolesService{}

type rolesService struct{}

func (r *rolesService) CreateRole(role *dto.CreateRoleParams) error {
	roleID := repositories.RoleRepositorysModules.CreateRole(role)
	if roleID == 0 {
		return errors.New("角色创建失败")
	}
	return r.CreateRolePermissionsRecord(&dto.CreateRolePermissionRecordParams{RoleID: strconv.FormatInt(roleID, 10), PageID: role.PageID})
}

func (r *rolesService) DeleteRole(id string) error {
	singleRoleResponse := repositories.RoleRepositorysModules.FindRoleById(id)
	if singleRoleResponse == nil {
		return errors.New("角色不存在")
	}
	return repositories.RoleRepositorysModules.DeleteRole(id)
}

func (r *rolesService) UpdateRole(role *dto.UpdateRoleParams, id string) error {
	singleRoleResponse := repositories.RoleRepositorysModules.FindRoleById(id)
	if singleRoleResponse == nil {
		return errors.New("角色不存在")
	}
	// 更新角色权限
	if role.PageID != nil {
		err := r.CreateRolePermissionsRecord(&dto.CreateRolePermissionRecordParams{RoleID: id, PageID: role.PageID})
		if err != nil {
			return err
		}
	}
	return repositories.RoleRepositorysModules.UpdateRole(role, id)

}
func (r *rolesService) GetRoles() ([]*dto.SingleRoleResponse, error) {
	return repositories.RoleRepositorysModules.GetRoles()
}

func (r *rolesService) CreateRolePermissionsRecord(params *dto.CreateRolePermissionRecordParams) error {
	// 检查角色是否存在
	if err := repositories.RoleRepositorysModules.CheckRolesExistence([]string{params.RoleID}); err != nil {
		return errors.New("角色不存在")
	}
	// 检查页面是否存在
	if params.PageID != nil {
		if err := repositories.PageRepositorysModules.CheckPagesExistence(params.PageID); err != nil {
			return errors.New("页面不存在")
		}
	}
	// 插入数据
	return repositories.RolesAndPagesRepository.CreateRecord(params)
}
