package monitor

import (
	"github.com/gin-gonic/gin"
	"vben/app/admin/model/monitor"
)

var (
	SysOperLog = serviceSysOperLog{}
)

type serviceSysOperLog struct{}

func (s serviceSysOperLog) AddSysOperLog(context *gin.Context, data string) {
	var operLog = monitor.SysOperLog{
		Title:      "错误日志",
		JsonResult: string(data),
		Status:     "1",
	}
	operLog.OperationLogAdd(context)
}
