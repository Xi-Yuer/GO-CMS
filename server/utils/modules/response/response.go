package responseModules

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Response = &response{}

type response struct {
}

type responseStruct struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *response) Success(context *gin.Context, v any) {
	context.JSON(http.StatusOK, responseStruct{
		Code: http.StatusOK,
		Data: v,
		Msg:  "成功",
	})
}

func (r *response) NoPermission(context *gin.Context, v any) {
	context.JSON(http.StatusOK, responseStruct{
		Code: http.StatusForbidden,
		Data: v,
		Msg:  "暂无权限",
	})
}

func (r *response) ServerError(context *gin.Context, v any) {
	context.JSON(http.StatusOK, responseStruct{
		Code: http.StatusInternalServerError,
		Data: v,
		Msg:  "服务器错误",
	})
}

func (r *response) NotFound(context *gin.Context, v any) {
	context.JSON(http.StatusOK, responseStruct{
		Code: http.StatusNotFound,
		Data: v,
		Msg:  "资源未找到",
	})
}

func (r *response) ParameterError(context *gin.Context, v any) {
	context.JSON(http.StatusOK, responseStruct{
		Code: http.StatusBadRequest,
		Data: v,
		Msg:  "参数错误",
	})
}

func (r *response) ParameterMissing(context *gin.Context, v any) {
	context.JSON(http.StatusOK, responseStruct{
		Code: http.StatusBadRequest,
		Data: v,
		Msg:  "参数缺失",
	})
}

func (r *response) ParameterTypeError(context *gin.Context, v any) {
	context.JSON(http.StatusOK, responseStruct{
		Code: http.StatusBadRequest,
		Data: v,
		Msg:  "参数类型错误",
	})
}

func (r *response) Conflict(context *gin.Context, v any) {
	context.JSON(http.StatusOK, responseStruct{
		Code: http.StatusConflict,
		Data: v,
		Msg:  "资源冲突",
	})
}

func (r *response) NoAuth(context *gin.Context, v any) {
	context.JSON(http.StatusOK, responseStruct{
		Code: http.StatusUnauthorized,
		Data: v,
		Msg:  "未登录",
	})
}
