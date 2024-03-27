package requestMiddlewareModule

import (
	"fmt"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
	"time"
)

var RequestMiddlewareModule = &requestMiddlewareModule{}

type requestMiddlewareModule struct{}

// RequestLoggerMiddleware  请求日志中间件
func (r *requestMiddlewareModule) RequestLoggerMiddleware(context *gin.Context) {
	path := context.Request.URL.Path
	methods := context.Request.Method
	params := context.Request.URL.RawQuery
	status := context.Writer.Status()
	start := time.Now()
	context.Next()
	duration := time.Since(start)
	utils.Log.Info(fmt.Sprintf("%v %s %s %s %s", status, methods, duration, path, params))
}
