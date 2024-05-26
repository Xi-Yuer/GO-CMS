package middlewares

import (
	authMiddleWareModule "github.com/Xi-Yuer/cms/middlewares/modules/auth"
	corsMiddlewareModule "github.com/Xi-Yuer/cms/middlewares/modules/cors"
	systemLogsMiddlewareModule "github.com/Xi-Yuer/cms/middlewares/modules/logs"
)
import sessionMiddleWareModule "github.com/Xi-Yuer/cms/middlewares/modules/session"

var SessionMiddleWareModule = sessionMiddleWareModule.Session

var AuthMiddleWareModule = authMiddleWareModule.AuthTokenMiddleWare

var LogsMiddlewareModule = systemLogsMiddlewareModule.SystemLogMiddlewareModule

var AuthMethodMiddleWare = authMiddleWareModule.AuthMethodMiddleWare

var AuthVerifyCookie = authMiddleWareModule.AuthVerifyCookie

var Cors = corsMiddlewareModule.Cors
