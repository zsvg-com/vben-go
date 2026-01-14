package group

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	// 配置相关
	g := e.Group("sys/group")
	{
		g.GET("", Get)
		g.GET("/info/:id", GetInfo)
		g.POST("", Post)
		g.PUT("", Put)
		g.DELETE("/:ids", Delete)
	}
}
