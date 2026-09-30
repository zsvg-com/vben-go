package oplog

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	g := e.Group("mon/oper")
	{
		g.GET("", Get)
		g.GET("/info/:id", GetInfo)
		g.POST("", Post)
		g.DELETE("/:ids", Delete)
	}
}
