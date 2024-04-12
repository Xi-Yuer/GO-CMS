package middlewares

import (
	authMiddleWareModule "github.com/Xi-Yuer/cms/middlewares/modules/auth"
	systemLogsMiddlewareModule "github.com/Xi-Yuer/cms/middlewares/modules/logs"
	requestMiddlewareModule "github.com/Xi-Yuer/cms/middlewares/modules/request"
)
import sessionMiddleWareModule "github.com/Xi-Yuer/cms/middlewares/modules/session"

var RequestMiddlewareModule = requestMiddlewareModule.RequestMiddlewareModule

var SessionMiddleWareModule = sessionMiddleWareModule.Session

var AuthMiddleWareModule = authMiddleWareModule.AuthMiddleWare

var LogsRepository = systemLogsMiddlewareModule.SystemLogMiddlewareModule
