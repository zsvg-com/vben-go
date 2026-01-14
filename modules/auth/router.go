package auth

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	// 配置相关
	g := e.Group("auth")
	{
		g.POST("/login", Login)
		g.POST("/logout", Logout)
		g.GET("/code", Code)
	}

	//u := e.Group("system/user")
	//{
	//	u.GET("/getInfo", GetInfo)
	//}
	e.GET("/system/user/getInfo", GetInfo)
	e.GET("/system/menu/getRouters", GetRouters)
}
