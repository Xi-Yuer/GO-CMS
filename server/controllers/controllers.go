package controllers

import (
	authControllersModules "github.com/Xi-Yuer/cms/controllers/modules/auth"
	userControllersModules "github.com/Xi-Yuer/cms/controllers/modules/users"
)

var UserController = userControllersModules.UserController
var AuthController = authControllersModules.AuthController
