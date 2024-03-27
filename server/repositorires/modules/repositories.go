package UserRepositoryModules

import (
	authRepositorysModules "github.com/Xi-Yuer/cms/repositorires/modules/auth"
	rolesRepositorysModules "github.com/Xi-Yuer/cms/repositorires/modules/roles"
	userRepositorysModules "github.com/Xi-Yuer/cms/repositorires/modules/users"
)

var UserRepositorysModules = userRepositorysModules.UserRepository
var AuthRepositorysModules = authRepositorysModules.AuthRepository
var RoleRepositorysModules = rolesRepositorysModules.RolesRepository
