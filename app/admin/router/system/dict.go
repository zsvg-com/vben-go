package system

import (
	"github.com/gin-gonic/gin"
	"vben/app/admin/api/system"
	"vben/app/core/utils/jwt"
)

func InitDict(e *gin.Engine) {
	v := e.Group("system/dict")
	{
		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			// 字典数据
			auth.GET("/data/list", system.ListDict)
			auth.POST("/data/export", system.ExportDict)
			auth.GET("/data/:dictCode", system.GetDictCode)
			auth.GET("/data/type/:dictType", system.GetDictDataByDictType)
			auth.POST("/data", system.SaveDictData)
			auth.PUT("/data", system.UpdateDictData)
			auth.DELETE("/data/:dictCodes", system.DeleteDictData)
			// 字典类型-字典管理
			auth.GET("/type/list", system.ListDictType)
			auth.POST("/type/export", system.ExportType)
			auth.GET("/type/:dictId", system.GetTypeDict)
			auth.POST("/type", system.SaveType)
			auth.PUT("/type", system.UpdateType)
			auth.DELETE("/type/:dictIds", system.DeleteDataType)
			auth.DELETE("/type/refreshCache", system.RefreshCache)
			auth.GET("/type/optionselect", system.GetOptionSelect)
		}
	}
}
