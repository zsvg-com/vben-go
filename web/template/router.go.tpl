package router

import (
	"github.com/gin-gonic/gin"
	 api "vben/app/admin/api/{{.ModuleName}}"
	"vben/app/pkg/utils"
)

func init{{.ClassName}}(e *gin.Engine) {
	// 路由权限相关
	v2 := e.Group("{{.ModuleName}}")
	{
		auth := v2.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			// 查询
            auth.GET("/{{.BusinessName}}/list", api.GetPageLimit{{.ClassName}})
            // 添加
            auth.POST("/{{.BusinessName}}", api.Create{{.ClassName}})
            // 修改
            auth.PUT("/{{.BusinessName}}", api.Update{{.ClassName}})
            // 删除
            auth.DELETE("/{{.BusinessName}}/:ids", api.Delete{{.ClassName}}ByIds)
            // 获取详情
            auth.GET("/{{.BusinessName}}/:id", api.Get{{.ClassName}})
		}
	}
}
