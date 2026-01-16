package singlec

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	g := e.Group("demo/single/cate")
	{
		g.GET("", Get)
		g.GET("/tree", GetTree)
		g.GET("/list", GetList)
		g.GET("/info/:id", GetInfo)
		g.POST("", Post)
		g.PUT("", Put)
		g.DELETE("/:ids", Delete)
	}
}
