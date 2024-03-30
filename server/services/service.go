package services

import (
	authServiceModules "github.com/Xi-Yuer/cms/services/modules/auth"
	pagesServiceModules "github.com/Xi-Yuer/cms/services/modules/pages"
	rolesServiceModules "github.com/Xi-Yuer/cms/services/modules/roles"
	userServiceModules "github.com/Xi-Yuer/cms/services/modules/users"
)

var UserService = userServiceModules.UserService
var AuthService = authServiceModules.AuthService
var RoleService = rolesServiceModules.RolesService
var PageService = pagesServiceModules.PageService
