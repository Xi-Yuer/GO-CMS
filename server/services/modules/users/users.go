package userServiceModules

import (
	"errors"
	"github.com/Xi-Yuer/cms/dto"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
	"github.com/Xi-Yuer/cms/utils"
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
	return repositories.UserRepositorysModules.CreateUser(user)
}

func (u *userService) FindUsersByAccount(account string) bool {
	return repositories.UserRepositorysModules.SelectUsersByAccount(account)
}

func (u *userService) FindUserById(id string) (*dto.UsersSingleResponse, bool) {
	return repositories.UserRepositorysModules.FindUserById(id)
}

func (u *userService) FindUserByParams(params *dto.QueryUsersParams) ([]dto.UsersSingleResponse, error) {
	return repositories.UserRepositorysModules.FindUserByParams(params)
}

func (u *userService) UpdateUser(params *dto.UpdateUserRequest, id string) error {
	_, exist := u.FindUserById(id)
	if !exist {
		return errors.New("用户不存在")
	}
	return repositories.UserRepositorysModules.UpdateUser(params, id)
}
func (u *userService) DeleteUser(id string) error {
	_, exist := u.FindUserById(id)
	if !exist {
		return errors.New("用户不存在")
	}
	return repositories.UserRepositorysModules.DeleteUser(id)
}

func (u *userService) FindUserByAccount(account string) (*dto.SingleUserResponseHasPassword, bool) {
	return repositories.UserRepositorysModules.FindUserByAccount(account)
}
