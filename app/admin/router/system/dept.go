package system

import (
	"github.com/gin-gonic/gin"
	"vben/app/admin/api/system"
	"vben/app/core/utils/jwt"
)

func InitDept(e *gin.Engine) {
	// 部门相关
	v := e.Group("system/dept")
	{
		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			// 部门数据
			auth.GET("/list", system.ListDept)
			auth.GET("/list/exclude/:deptId", system.ExcludeDept)
			auth.GET("/:deptId", system.GetDept)
			auth.POST("", system.SaveDept)
			auth.PUT("", system.UpDataDept)
			auth.DELETE("/:deptId", system.DeleteDept)
		}
	}
}
