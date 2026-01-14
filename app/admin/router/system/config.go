package system

import (
	"github.com/gin-gonic/gin"
	"vben/app/admin/api/system"
	"vben/app/core/utils/jwt"
)

func InitConfig(e *gin.Engine) {
	// 配置相关
	v := e.Group("system/config")
	{
		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.GET("/list", system.ListConfig)
			auth.POST("/export", system.ExportConfig)
			auth.GET("/:configId", system.GetConfigInfo)
			auth.GET("/configKey/:configKey", system.GetConfigKey)
			auth.POST("", system.SaveConfig)
			auth.PUT("", system.UploadConfig)
			auth.DELETE("/:configIds", system.DetectConfig)
			auth.DELETE("/refreshCache", system.RefreshCacheConfig)
		}
	}
}
