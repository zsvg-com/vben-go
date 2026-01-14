package monitor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vben/app/admin/model/monitor"
)

func ServerData(context *gin.Context) {
	var server = monitor.GetServerInfo(context)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": server,
	})
}
