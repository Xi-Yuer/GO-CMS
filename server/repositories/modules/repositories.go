package UserRepositoryModules

import (
	authRepositorysModules "github.com/Xi-Yuer/cms/repositories/modules/auth"
	rolesRepositorysModules "github.com/Xi-Yuer/cms/repositories/modules/roles"
	userRepositorysModules "github.com/Xi-Yuer/cms/repositories/modules/users"
)

var UserRepositorysModules = userRepositorysModules.UserRepository
var AuthRepositorysModules = authRepositorysModules.AuthRepository
var RoleRepositorysModules = rolesRepositorysModules.RolesRepository
