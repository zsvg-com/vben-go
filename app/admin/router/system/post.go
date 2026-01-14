package system

import (
	"github.com/gin-gonic/gin"
	"vben/app/admin/api/system"
	"vben/app/core/utils/jwt"
)

func InitPost(e *gin.Engine) {
	v := e.Group("system/post")
	{
		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.GET("/list", system.ListPost)
			auth.POST("/export", system.ExportPost)
			auth.GET("/:postId", system.GetPostInfo)
			auth.POST("", system.SavePost)
			auth.PUT("", system.UploadPost)
			auth.DELETE("/:postIds", system.DeletePost)
			auth.GET("/optionselect", system.GetPostOptionSelect)
		}
	}
}
