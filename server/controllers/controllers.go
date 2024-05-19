package controllers

import (
	authControllersModules "github.com/Xi-Yuer/cms/controllers/modules/auth"
	commitsControllerModules "github.com/Xi-Yuer/cms/controllers/modules/commits"
	departmentControllerModules "github.com/Xi-Yuer/cms/controllers/modules/department"
	interfaceControllerModules "github.com/Xi-Yuer/cms/controllers/modules/interface"
	logsControllerModules "github.com/Xi-Yuer/cms/controllers/modules/logs"
	pagesControllerModules "github.com/Xi-Yuer/cms/controllers/modules/pages"
	roleControllersModules "github.com/Xi-Yuer/cms/controllers/modules/role"
	systemControllerModules "github.com/Xi-Yuer/cms/controllers/modules/system"
	templateControllerModules "github.com/Xi-Yuer/cms/controllers/modules/template"
	"github.com/Xi-Yuer/cms/controllers/modules/timeTask"
	uploadControllerModules "github.com/Xi-Yuer/cms/controllers/modules/upload"
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
var SystemController = systemControllerModules.SystemController
var TimeTaskController = timeTaskControllerModules.TimeTaskController
var UploadController = uploadControllerModules.UploadController
var TemplateController = templateControllerModules.TemplateController
