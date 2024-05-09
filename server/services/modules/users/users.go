package userServiceModules

import (
	"errors"
	"github.com/Xi-Yuer/cms/dto"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
	usersAndRolesServiceModules "github.com/Xi-Yuer/cms/services/modules/usersAndRoles"
	"github.com/Xi-Yuer/cms/utils"
	"strconv"
)

var UserService = &userService{}

type userService struct{}

func (u *userService) GetUser(id string) (*dto.UsersSingleResponse, error) {
	return repositories.UserRepositorysModules.GetUser(id)
}

func (u *userService) GetUsers(page dto.Page) ([]dto.UsersSingleResponse, error) {
	return repositories.UserRepositorysModules.GetUsers(page)
}

func (u *userService) CreateUser(user *dto.CreateSingleUserRequest) error {
	exist := u.FindUsersByAccount(user.Account)
	if exist {
		return errors.New("用户名已存在")
	}
	password, err := utils.Bcrypt.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = password
	id := repositories.UserRepositorysModules.CreateUser(user)
	if id == 0 {
		return errors.New("创建用户失败")
	}
	// 给用户分配角色信息
	if user.RoleID != nil {
		if err = repositories.UsersAndRolesRepositorys.CreateRecords(strconv.FormatInt(id, 10), user.RoleID); err != nil {
			return err
		}
	}
	return err
}

func (u *userService) FindUsersByAccount(account string) bool {
	return repositories.UserRepositorysModules.SelectUsersByAccount(account)
}

func (u *userService) FindUserById(id string) (*dto.UsersSingleResponse, bool) {
	return repositories.UserRepositorysModules.FindUserById(id)
}

func (u *userService) FindUserByParams(params *dto.QueryUsersParams) (*dto.HasTotalResponseData, error) {
	return repositories.UserRepositorysModules.FindUserByParams(params)
}
func (u *userService) FindUserByParamsAndOutRoleID(id string, params *dto.QueryUsersParams) (*dto.HasTotalResponseData, error) {
	return repositories.UserRepositorysModules.FindUserByParamsAndOutRoleID(id, params)
}

func (u *userService) UpdateUser(params *dto.UpdateUserRequest, id string) error {
	user, exist := u.FindUserById(id)
	if !exist {
		return errors.New("资源不存在")
	}
	// 超级管理员无法修改
	if user.IsAdmin == 1 {
		return errors.New("系统账号禁止修改")
	}
	if params.RolesID != nil && len(params.RolesID) > 0 {
		// 给用户分配角色信息
		err := repositories.UsersAndRolesRepositorys.CreateRecords(id, params.RolesID)
		if err != nil {
			return err
		}
	}
	return repositories.UserRepositorysModules.UpdateUser(params, id)
}

func (u *userService) DeleteUser(id string) error {
	user, exist := u.FindUserById(id)
	if !exist {
		return errors.New("资源不存在")
	}
	if user.IsAdmin == 1 {
		return errors.New("系统账号禁止删除")
	}
	return repositories.UserRepositorysModules.DeleteUser(id)
}

func (u *userService) FindUserByAccount(account string) (*dto.SingleUserResponseHasPassword, bool) {
	return repositories.UserRepositorysModules.FindUserByAccount(account)
}

func (u *userService) GetUserByRoleID(roleID string, params dto.Page) (*dto.HasTotalResponseData, error) {
	if err := usersAndRolesServiceModules.UserAndRolesService.FindRoleById(roleID); err != nil {
		return nil, err
	}
	return repositories.UserRepositorysModules.GetUserByRoleID(roleID, params)
}

func (u *userService) ExportExcel(IDs *dto.ExportExcelResponse) ([]*dto.UsersSingleResponse, error) {
	return repositories.UserRepositorysModules.ExportExcel(IDs)
}
