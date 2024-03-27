package middlewares

import (
	authMiddleWareModule "github.com/Xi-Yuer/cms/middlewares/modules/auth"
	requestMiddlewareModule "github.com/Xi-Yuer/cms/middlewares/modules/request"
)
import sessionMiddleWareModule "github.com/Xi-Yuer/cms/middlewares/modules/session"

var RequestMiddlewareModule = requestMiddlewareModule.RequestMiddlewareModule

var SessionMiddleWareModule = sessionMiddleWareModule.Session

var AuthMiddleWareModule = authMiddleWareModule.AuthMiddleWare
