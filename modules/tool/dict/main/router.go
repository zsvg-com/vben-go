package dict

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	g := e.Group("tool/dict/main")
	{
		g.GET("", Get)
		g.GET("/list", GetList)
		g.GET("/info/:id", GetInfo)
		g.POST("", Post)
		g.PUT("", Put)
		g.DELETE("/:ids", Delete)
	}
}
