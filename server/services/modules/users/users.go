package userServiceModules

import (
	repositories "github.com/Xi-Yuer/cms/repositorires/modules"
	"github.com/Xi-Yuer/cms/responsies"
)

var UserService = &userService{}

type userService struct{}

func (u *userService) GetUser(id string) (*responsies.UsersSingleResponse, error) {
	return repositories.UserRepositorysModules.GetUser(id)
}
func (u *userService) GetUsers() (*[]responsies.UsersSingleResponse, error) {
	return nil, nil
}

func (u *userService) CreateUser(user *responsies.CreateSingleUserRequest) error {
	return repositories.UserRepositorysModules.CreateUser(user)
}

func (u *userService) UpdateUser() {

}

func (u *userService) DeleteUser() {

}
