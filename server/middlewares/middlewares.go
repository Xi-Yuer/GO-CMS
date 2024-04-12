package middlewares

import (
	authMiddleWareModule "github.com/Xi-Yuer/cms/middlewares/modules/auth"
	systemLogsMiddlewareModule "github.com/Xi-Yuer/cms/middlewares/modules/logs"
)
import sessionMiddleWareModule "github.com/Xi-Yuer/cms/middlewares/modules/session"

var SessionMiddleWareModule = sessionMiddleWareModule.Session

var AuthMiddleWareModule = authMiddleWareModule.AuthMiddleWare

var LogsMiddlewareModule = systemLogsMiddlewareModule.SystemLogMiddlewareModule
