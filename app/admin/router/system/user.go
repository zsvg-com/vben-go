package system

import (
	"github.com/gin-gonic/gin"
	"vben/app/admin/api/system"
	"vben/app/core/utils/jwt"
)

func InitUser(e *gin.Engine) {
	v := e.Group("system")
	{
		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.GET("/user/list", system.ListUser)
			auth.POST("/user/export", system.ExportExport)
			auth.POST("/user/importData", system.ImportUserData)
			auth.POST("/user/importTemplate", system.ImportTemplate)
			auth.GET("/user/:userId", system.GetUserInfo)
			auth.GET("/user/", system.GetUserInfo)
			auth.POST("/user", system.SaveUser)
			auth.PUT("/user", system.UploadUser)
			// 更新用户信息 验证防重复提交和限流中间件
			//auth
			//.Use(utils.RepeatSubmitMiddleware(500))
			//.Use(utils.RateLimiterMiddleware(60, 20, utils.DEFAULT))
			//.PUT("/user", system.UploadUser)
			auth.DELETE("/user/:userIds", system.DeleteUserById)
			auth.PUT("/user/resetPwd", system.ResetPwd)
			auth.PUT("/user/changeStatus", system.ChangeUserStatus)
			auth.GET("/user/authRole/:userId", system.GetAuthUserRole)
			auth.PUT("/user/authRole", system.PutAuthUser)
			auth.GET("/user/deptTree", system.GetUserDeptTree)
		}
	}
}
