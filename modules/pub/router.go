package pub

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	// 配置相关
	g := e.Group("pub")
	{
		g.GET("/actor", GetActor)
		g.POST("/rece", PostRece)
		g.GET("/rece", GetRece)
		g.GET("/group/tree", GetGroupTree)
		g.GET("/user", GetUser)
		g.GET("/menus", GetMenus)
	}
}
