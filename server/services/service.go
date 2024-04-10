package services

import (
	authServiceModules "github.com/Xi-Yuer/cms/services/modules/auth"
	departmentServiceModules "github.com/Xi-Yuer/cms/services/modules/department"
	interfaceServiceModules "github.com/Xi-Yuer/cms/services/modules/interface"
	pagesServiceModules "github.com/Xi-Yuer/cms/services/modules/pages"
	rolesServiceModules "github.com/Xi-Yuer/cms/services/modules/roles"
	userServiceModules "github.com/Xi-Yuer/cms/services/modules/users"
	usersAndRolesServiceModules "github.com/Xi-Yuer/cms/services/modules/usersAndRoles"
)

var UserService = userServiceModules.UserService
var AuthService = authServiceModules.AuthService
var RoleService = rolesServiceModules.RolesService
var PageService = pagesServiceModules.PageService
var UserAndRolesService = usersAndRolesServiceModules.UserAndRolesService
var DepartmentService = departmentServiceModules.DepartmentService
var InterfaceService = interfaceServiceModules.InterfaceService
