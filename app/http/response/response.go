package response

import (
	"github.com/gin-gonic/gin"
	"goinspect.cn/app/http"
)

type Response struct {
	Code	int 		`json:"code"`
	Message string		`json:"msg"`
	Data	interface{}	`json:"data"`
}

func Json(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(200, Response{
		Code: code,
		Message: msg,
		Data: data,
	})
	return
}

func Success(ctx *gin.Context, data interface{}) {
	Json(ctx, http.CODE_SUCCESS, "请求成功", data)
	return
}

func Error(ctx *gin.Context, code int, msg string) {
	Json(ctx, code, msg, nil)
	return
}