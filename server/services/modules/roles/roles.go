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
	return r.CreateRolePermissionsRecord(&dto.CreateRolePermissionRecordParams{RoleID: strconv.FormatInt(roleID, 10), PageID: role.PageID, InterfaceID: role.InterfaceID})
}

func (r *rolesService) DeleteRole(id string) error {
	singleRoleResponse := repositories.RoleRepositorysModules.FindRoleById(id)
	if singleRoleResponse == nil {
		return errors.New("资源不存在")
	}
	if singleRoleResponse.CanEdit == 0 {
		return errors.New("该角色禁止删除")
	}
	return repositories.RoleRepositorysModules.DeleteRole(id)
}

func (r *rolesService) UpdateRole(role *dto.UpdateRoleParams, id string) error {
	singleRoleResponse := repositories.RoleRepositorysModules.FindRoleById(id)
	if singleRoleResponse == nil {
		return errors.New("资源不存在")
	}
	if singleRoleResponse.CanEdit == 0 {
		return errors.New("该角色禁止编辑")
	}
	var err error
	// 更新角色权限
	err = r.CreateRolePermissionsRecord(&dto.CreateRolePermissionRecordParams{RoleID: id, PageID: role.PageID, InterfaceID: role.InterfaceID})
	if err != nil {
		return err
	}
	// 更新角色基本信息
	err = repositories.RoleRepositorysModules.UpdateRole(role, id)
	return err

}
func (r *rolesService) GetRoles(params *dto.QueryRolesParams) (*dto.HasTotalResponseData, error) {
	return repositories.RoleRepositorysModules.GetRoles(params)
}

func (r *rolesService) CreateRolePermissionsRecord(params *dto.CreateRolePermissionRecordParams) error {
	// 检查角色是否存在
	if err := repositories.RoleRepositorysModules.CheckRolesExistence([]string{params.RoleID}); err != nil {
		return errors.New("资源不存在")
	}
	// 检查页面是否存在
	if params.PageID != nil && len(params.PageID) > 0 {
		if err := repositories.PageRepositorysModules.CheckPagesExistence(params.PageID); err != nil {
			return errors.New("资源不存在")
		}
	}

	if params.InterfaceID != nil && len(params.InterfaceID) > 0 {
		if err := repositories.InterfaceRepository.CheckInterfacesExistence(params.InterfaceID); err != nil {
			return errors.New("资源不存在")
		}
	}
	var err error
	// 创建角色接口权限
	err = repositories.RolesAndInterfacesRepository.CreateRecord(params)
	if err != nil {
		return err
	}
	// 创建角色页面权限
	err = repositories.RolesAndPagesRepository.CreateRecord(params)
	return err
}
func (r *rolesService) ExportExcel(params *dto.ExportExcelResponse) ([]*dto.SingleRoleResponse, error) {
	return repositories.RoleRepositorysModules.ExportExcel(params)
}

func (r *rolesService) CreateOneRecord(params *dto.CreateOneRecord) error {
	return repositories.UsersAndRolesRepositorys.CreateOneRecord(params)
}

func (r *rolesService) DeleteOneRecord(params *dto.DeleteOneRecord) error {
	user, b := repositories.UserRepositorysModules.FindUserById(params.UserID)
	if !b {
		return errors.New("资源不存在")
	}
	if user.IsAdmin == 1 {
		return errors.New("系统资源，禁止删除")
	}
	return repositories.UsersAndRolesRepositorys.DeleteOneRecord(params)
}
