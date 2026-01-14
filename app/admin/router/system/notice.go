package system

import (
	"github.com/gin-gonic/gin"
	"vben/app/admin/api/system"
	"vben/app/core/utils/jwt"
)

func InitNotice(e *gin.Engine) {
	// 消息相关
	v := e.Group("system/notice")
	{
		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.GET("/list", system.ListNotice)
			auth.GET("/:noticeId", system.GetNotice)
			auth.POST("", system.SaveNotice)
			auth.PUT("", system.UploadNotice)
			auth.DELETE("/:noticeIds", system.DeleteNotice)
		}
	}
}
