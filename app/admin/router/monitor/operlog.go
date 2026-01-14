package monitor

import (
	"github.com/gin-gonic/gin"
	"vben/app/admin/api/monitor"
	"vben/app/core/utils/jwt"
)

func InitOperlog(e *gin.Engine) {
	// 操作日志相关
	v := e.Group("monitor/operlog")
	{
		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.GET("/list", monitor.ListOperlog)
			auth.DELETE("/:operId", monitor.DelectOperlog)
			auth.DELETE("/clean", monitor.ClearOperlog)
			auth.POST("/export", monitor.ExportOperlog)
		}
	}
}
