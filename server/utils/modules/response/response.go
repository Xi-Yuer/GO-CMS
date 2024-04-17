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

func (r *response) FileExist(context *gin.Context, v any) {
	context.JSON(http.StatusOK, responseStruct{
		Code: http.StatusContinue,
		Data: v,
		Msg:  "文件不存在，继续上传文件",
	})
}

func (r *response) NoPermission(context *gin.Context, v string) {
	context.JSON(http.StatusForbidden, responseStruct{
		Code: http.StatusForbidden,
		Data: nil,
		Msg:  v,
	})
}

func (r *response) ServerError(context *gin.Context, v string) {
	context.JSON(http.StatusInternalServerError, responseStruct{
		Code: http.StatusInternalServerError,
		Data: nil,
		Msg:  v,
	})
}

func (r *response) NotFound(context *gin.Context, v string) {
	context.JSON(http.StatusNotFound, responseStruct{
		Code: http.StatusNotFound,
		Msg:  v,
		Data: nil,
	})
}

func (r *response) ParameterError(context *gin.Context, v string) {
	context.JSON(http.StatusBadRequest, responseStruct{
		Code: http.StatusBadRequest,
		Data: nil,
		Msg:  v,
	})
}

func (r *response) ParameterMissing(context *gin.Context, v string) {
	context.JSON(http.StatusBadRequest, responseStruct{
		Code: http.StatusBadRequest,
		Data: nil,
		Msg:  v,
	})
}

func (r *response) ParameterTypeError(context *gin.Context, v string) {
	context.JSON(http.StatusBadRequest, responseStruct{
		Code: http.StatusBadRequest,
		Data: nil,
		Msg:  v,
	})
}

func (r *response) Conflict(context *gin.Context, v string) {
	context.JSON(http.StatusConflict, responseStruct{
		Code: http.StatusConflict,
		Data: nil,
		Msg:  v,
	})
}

func (r *response) NoAuth(context *gin.Context, v string) {
	context.JSON(http.StatusOK, responseStruct{
		Code: http.StatusUnauthorized,
		Data: nil,
		Msg:  v,
	})
}
