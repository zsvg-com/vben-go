package monitor

import (
	"github.com/gin-gonic/gin"
	"vben/app/admin/api/monitor"
	"vben/app/core/utils/jwt"
)

func InitServer(e *gin.Engine) {
	// 服务监控
	v := e.Group("monitor")
	{
		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.GET("/server", monitor.ServerData)
		}
	}
}
