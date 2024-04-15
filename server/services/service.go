package services

import (
	authServiceModules "github.com/Xi-Yuer/cms/services/modules/auth"
	departmentServiceModules "github.com/Xi-Yuer/cms/services/modules/department"
	interfaceServiceModules "github.com/Xi-Yuer/cms/services/modules/interface"
	logsServiceModules "github.com/Xi-Yuer/cms/services/modules/logs"
	pagesServiceModules "github.com/Xi-Yuer/cms/services/modules/pages"
	rolesServiceModules "github.com/Xi-Yuer/cms/services/modules/roles"
	rolesAndInterfacesServiceModules "github.com/Xi-Yuer/cms/services/modules/rolesAndInterfaces"
	timeTaskServiceModules "github.com/Xi-Yuer/cms/services/modules/timeTask"
	userServiceModules "github.com/Xi-Yuer/cms/services/modules/users"
)

var UserService = userServiceModules.UserService
var AuthService = authServiceModules.AuthService
var RoleService = rolesServiceModules.RolesService
var PageService = pagesServiceModules.PageService
var DepartmentService = departmentServiceModules.DepartmentService
var InterfaceService = interfaceServiceModules.InterfaceService
var LogsService = logsServiceModules.LogsService
var RolesAndInterfacesService = rolesAndInterfacesServiceModules.RolesAndInterfacesService
var TimeTaskService = timeTaskServiceModules.TimeTaskService
