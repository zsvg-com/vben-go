package system

import (
	"github.com/gin-gonic/gin"
	"vben/app/admin/api/system"
	"vben/app/core/utils/jwt"
)

func InitRole(e *gin.Engine) {
	v := e.Group("system/role")
	{
		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.GET("/list", system.ListRole)
			auth.POST("/export", system.ExportRole)
			auth.GET("/:roleId", system.GetRoleInfo)
			auth.POST("", system.SaveRole)
			auth.PUT("", system.UploadRole)
			auth.PUT("/dataScope", system.PutDataScope)
			auth.PUT("/changeStatus", system.ChangeRoleStatus)
			auth.DELETE("/:roleIds", system.DeleteRole)
			auth.GET("/optionselect", system.GetRoleOptionSelect)
			auth.GET("/authUser/allocatedList", system.GetAllocatedList)
			auth.GET("/authUser/unallocatedList", system.GetUnAllocatedList)
			auth.PUT("/authUser/cancel", system.CancelRole)
			auth.PUT("/authUser/cancelAll", system.CancelAllRole)
			auth.PUT("/authUser/selectAll", system.SelectRoleAll)
			auth.GET("/deptTree/:roleId", system.GetDeptTreeRole)
		}
	}
}
