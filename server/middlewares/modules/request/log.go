package requestMiddlewareModule

import (
	"fmt"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
)

var RequestMiddlewareModule = &requestMiddlewareModule{}

type requestMiddlewareModule struct{}

// RequestLoggerMiddleware  请求日志中间件
func (r *requestMiddlewareModule) RequestLoggerMiddleware(context *gin.Context) {

	path := context.Request.URL.Path
	methods := context.Request.Method
	params := context.Request.URL.RawQuery
	body := context.Request.Body

	utils.Log.Info(fmt.Sprintf("%s %s %s %s", methods, path, params, body))
	context.Next()
}
