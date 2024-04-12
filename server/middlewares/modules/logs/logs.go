package systemLogsMiddlewareModule

import (
	"github.com/Xi-Yuer/cms/constant"
	"github.com/Xi-Yuer/cms/dto"
	repositories "github.com/Xi-Yuer/cms/repositories/modules"
	"github.com/Xi-Yuer/cms/utils"
	"github.com/gin-gonic/gin"
	"time"
)

var SystemLogMiddlewareModule = &systemLogMiddlewareModule{}

type systemLogMiddlewareModule struct{}

// SystemLogMiddleware 系统日志中间件
// 用于记录系统日志的中间件
// 例如：记录请求的URL、请求的方法、请求的参数、请求的IP地址等
// 还可以记录响应的状态码、响应的时间等
// 可以根据需要记录更多的信息
// 该中间件可以用于记录系统日志，以便进行问题排查和分析
func (m *systemLogMiddlewareModule) SystemLogMiddleware(context *gin.Context) {
	path := context.Request.URL.Path
	methods := context.Request.Method
	status := context.Writer.Status()
	ip := context.ClientIP()
	start := time.Now()
	context.Next()
	duration := time.Since(start)
	var user *dto.JWTPayload
	value, exists := context.Get(constant.JWTPAYLOAD)
	if !exists {
		user.NickName = "未登录用户"
		user.ID = ""
	} else {
		user = value.(*dto.JWTPayload)
	}
	params := dto.CreateLogRecordRequest{
		ID:              utils.GenID(),
		UserName:        user.NickName,
		UserID:          user.ID,
		UserIP:          ip,
		RequestMethod:   methods,
		RequestPath:     path,
		RequestStatus:   status,
		RequestDuration: duration.String(),
	}
	go func() {
		_ = repositories.LogsRepository.CreateLogRecord(&params)
	}()

}
