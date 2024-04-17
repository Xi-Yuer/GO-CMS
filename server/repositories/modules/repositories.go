package repositories

import (
	commitsRepositoryModules "github.com/Xi-Yuer/cms/repositories/modules/commits"
	departmentRepositoryModules "github.com/Xi-Yuer/cms/repositories/modules/department"
	interfaceRepositoryModules "github.com/Xi-Yuer/cms/repositories/modules/interface"
	logsRepositoryModules "github.com/Xi-Yuer/cms/repositories/modules/logs"
	pagesRepositoryModules "github.com/Xi-Yuer/cms/repositories/modules/pages"
	rolesRepositorysModules "github.com/Xi-Yuer/cms/repositories/modules/roles"
	rolesAndInterfacesRepositoryModules "github.com/Xi-Yuer/cms/repositories/modules/rolesAndInterfaces"
	rolesAndPagesRepositoryModules "github.com/Xi-Yuer/cms/repositories/modules/rolesAndPages"
	uploadRepositorysModules "github.com/Xi-Yuer/cms/repositories/modules/upload"
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
var LogsRepository = logsRepositoryModules.LogsRepository
var RolesAndInterfacesRepository = rolesAndInterfacesRepositoryModules.RolesAndInterfacesRepository
var CommitsRepositoryModules = commitsRepositoryModules.CommitsRepositoryModules
var UploadRepository = uploadRepositorysModules.UploadRepository
