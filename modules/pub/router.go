package pub

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	// 配置相关
	g := e.Group("pub")
	{
		g.GET("/org", GetOrg)
		g.POST("/org/rece", PostRece)
		g.GET("/org/rece", GetRece)
		g.GET("/org/group/tree", GetGroupTree)
	}

}
