package R

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type R struct {
	Ctx *gin.Context
}

// 返回的结果：
type Result struct {
	Code int         `json:"code"`           //提示代码
	Msg  string      `json:"msg"`            //提示信息
	Data interface{} `json:"data,omitempty"` //数据
}

func ReturnSuccess(data any) Result {
	return Result{
		Code: http.StatusOK,
		Msg:  "操作成功",
		Data: data,
	}
}

func ReturnSuccessMsg(msg string) Result {
	return Result{
		Code: http.StatusOK,
		Msg:  msg,
	}
}

func ReturnFailMsg(msg string) Result {
	return Result{
		Msg:  msg,
		Code: http.StatusInternalServerError,
	}
}
