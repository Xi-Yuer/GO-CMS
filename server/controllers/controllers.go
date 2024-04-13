package controllers

import (
	authControllersModules "github.com/Xi-Yuer/cms/controllers/modules/auth"
	commitsControllerModules "github.com/Xi-Yuer/cms/controllers/modules/commits"
	departmentControllerModules "github.com/Xi-Yuer/cms/controllers/modules/department"
	interfaceControllerModules "github.com/Xi-Yuer/cms/controllers/modules/interface"
	logsControllerModules "github.com/Xi-Yuer/cms/controllers/modules/logs"
	pagesControllerModules "github.com/Xi-Yuer/cms/controllers/modules/pages"
	roleControllersModules "github.com/Xi-Yuer/cms/controllers/modules/role"
	userControllersModules "github.com/Xi-Yuer/cms/controllers/modules/users"
)

var UserController = userControllersModules.UserController
var AuthController = authControllersModules.AuthController
var RoleController = roleControllersModules.RoleController
var PagesController = pagesControllerModules.PagesController
var DepartmentController = departmentControllerModules.DepartmentController
var InterfaceController = interfaceControllerModules.InterfaceController
var LogsController = logsControllerModules.LogsController
var CommitsController = commitsControllerModules.CommitsController
