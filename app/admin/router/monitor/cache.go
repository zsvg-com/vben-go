package monitor

import (
	"github.com/gin-gonic/gin"
	"vben/app/admin/api/monitor"
	"vben/app/core/utils/jwt"
)

func InitCache(e *gin.Engine) {
	// 缓存相关
	v := e.Group("monitor/cache")
	{
		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.GET("", monitor.CacheHandler)
			auth.GET("getNames", monitor.CacheHandler)
			auth.GET("getKeys/:cacheName", monitor.GetCacheKeysHandler)
			auth.GET("getValue/:cacheName/:cacheKey", monitor.GetCacheValueHandler)
			auth.DELETE("clearCacheName/:cacheName", monitor.ClearCacheNameHandler)
			auth.DELETE("clearCacheKey/:cacheKey", monitor.ClearCacheKeyHandler)
			auth.DELETE("clearCacheAll", monitor.ClearCacheAllHandler)
		}
	}
}
