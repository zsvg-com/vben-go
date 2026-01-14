package role

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	g := e.Group("sys/role")
	{
		g.GET("", Get)
		g.GET("/info/:id", GetInfo)
		g.POST("", Post)
		g.PUT("", Put)
		g.DELETE("/:ids", Delete)
		g.GET("/perms", GetPerms)
	}
}
