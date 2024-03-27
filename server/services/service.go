package services

import (
	authServiceModules "github.com/Xi-Yuer/cms/services/modules/auth"
	userServiceModules "github.com/Xi-Yuer/cms/services/modules/users"
)

var UserService = userServiceModules.UserService
var AuthService = authServiceModules.AuthService
