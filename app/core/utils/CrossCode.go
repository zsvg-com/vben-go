package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 解决跨域问题
func Core() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token")
		context.Header("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE,OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		context.Header("Access-Control-Allow-Credentials", "True")
		//放行索引options
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		//处理请求
		context.Next()
	}
}
