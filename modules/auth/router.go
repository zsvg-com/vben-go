package auth

import (
	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	g := e.Group("auth")
	{
		g.POST("/login", Login)
		g.POST("/logout", Logout)
		g.GET("/code", Code)
	}

}
