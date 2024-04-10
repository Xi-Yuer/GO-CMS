package controllers

import (
	authControllersModules "github.com/Xi-Yuer/cms/controllers/modules/auth"
	departmentControllerModules "github.com/Xi-Yuer/cms/controllers/modules/department"
	interfaceControllerModules "github.com/Xi-Yuer/cms/controllers/modules/interface"
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
