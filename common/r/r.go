package r

import "net/http"

type R struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Ok(data any) R {
	return R{
		Msg:  "成功",
		Code: http.StatusOK,
		Data: data,
	}
}
