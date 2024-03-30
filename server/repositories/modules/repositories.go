package UserRepositoryModules

import (
	authRepositorysModules "github.com/Xi-Yuer/cms/repositories/modules/auth"
	rolesRepositorysModules "github.com/Xi-Yuer/cms/repositories/modules/roles"
	userRepositorysModules "github.com/Xi-Yuer/cms/repositories/modules/users"
	usersAndRolesRepositorysModules "github.com/Xi-Yuer/cms/repositories/modules/usersAndRoles"
)

var UserRepositorysModules = userRepositorysModules.UserRepository
var AuthRepositorysModules = authRepositorysModules.AuthRepository
var RoleRepositorysModules = rolesRepositorysModules.RolesRepository
var UsersAndRolesRepositorys = usersAndRolesRepositorysModules.UsersAndRolesRepositorys
