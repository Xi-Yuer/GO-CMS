package controllers

import (
	authControllersModules "github.com/Xi-Yuer/cms/controllers/modules/auth"
	roleControllersModules "github.com/Xi-Yuer/cms/controllers/modules/role"
	userControllersModules "github.com/Xi-Yuer/cms/controllers/modules/users"
)

var UserController = userControllersModules.UserController
var AuthController = authControllersModules.AuthController
var RoleController = roleControllersModules.RoleController
