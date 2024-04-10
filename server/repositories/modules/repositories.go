package repositories

import (
	departmentRepositoryModules "github.com/Xi-Yuer/cms/repositories/modules/department"
	interfaceRepositoryModules "github.com/Xi-Yuer/cms/repositories/modules/interface"
	pagesRepositoryModules "github.com/Xi-Yuer/cms/repositories/modules/pages"
	rolesRepositorysModules "github.com/Xi-Yuer/cms/repositories/modules/roles"
	rolesAndPagesRepositoryModules "github.com/Xi-Yuer/cms/repositories/modules/rolesAndPages"
	userRepositorysModules "github.com/Xi-Yuer/cms/repositories/modules/users"
	usersAndRolesRepositorysModules "github.com/Xi-Yuer/cms/repositories/modules/usersAndRoles"
)

var UserRepositorysModules = userRepositorysModules.UserRepository
var RoleRepositorysModules = rolesRepositorysModules.RolesRepository
var UsersAndRolesRepositorys = usersAndRolesRepositorysModules.UsersAndRolesRepositorys
var PageRepositorysModules = pagesRepositoryModules.PageRepository
var RolesAndPagesRepository = rolesAndPagesRepositoryModules.RolesAndPagesRepository
var DepartmentRepository = departmentRepositoryModules.DepartmentRepository
var InterfaceRepository = interfaceRepositoryModules.InterfaceRepository
