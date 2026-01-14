package monitor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vben/app/admin/model/constants"
	"vben/app/admin/model/monitor"
	"vben/pkg/cache/redisCache"
)

func CacheHandler(context *gin.Context) {
	var list []monitor.SysCache
	list = append(list, monitor.SysCache{
		CacheName: constants.LoginCacheKey,
		Remark:    "用户信息",
	})
	list = append(list, monitor.SysCache{
		CacheName: "sys_config:",
		Remark:    "配置信息",
	})
	list = append(list, monitor.SysCache{
		CacheName: constants.SysDictCacheKey,
		Remark:    "数据字典",
	})
	list = append(list, monitor.SysCache{
		CacheName: "captcha_codes:",
		Remark:    "验证码",
	})
	list = append(list, monitor.SysCache{
		CacheName: "repeat_submit:",
		Remark:    "防重提交",
	})
	list = append(list, monitor.SysCache{
		CacheName: "rate_limit:",
		Remark:    "限流处理",
	})
	list = append(list, monitor.SysCache{
		CacheName: "pwd_err_cnt:",
		Remark:    "密码错误次数",
	})
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": list,
	})
}

func GetCacheKeysHandler(context *gin.Context) {
	cacheName := context.Param("cacheName")
	keys, _, err := redisCache.NewRedisCache().Scan(0, cacheName+"*", constants.ScanCountMax)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": http.StatusInternalServerError,
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": keys,
	})
}

func GetCacheValueHandler(context *gin.Context) {
	cacheName := context.Param("cacheName")
	cacheKey := context.Param("cacheKey")
	value, err := redisCache.NewRedisCache().Get(cacheKey)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": http.StatusInternalServerError,
		})
	}
	var cache = monitor.SysCache{
		CacheName:  cacheName,
		CacheKey:   cacheKey,
		CacheValue: value,
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": cache,
	})
}

func ClearCacheNameHandler(context *gin.Context) {
	cacheName := context.Param("cacheName")
	keys, _, err := redisCache.NewRedisCache().Scan(0, cacheName+"*", constants.ScanCountMax)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": http.StatusInternalServerError,
		})
	}
	for i := range keys {
		redisCache.NewRedisCache().Del(keys[i])
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
	})
}

func ClearCacheKeyHandler(context *gin.Context) {
	cacheKey := context.Param("cacheKey")
	_, err := redisCache.NewRedisCache().Del(cacheKey)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": http.StatusInternalServerError,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
	})
}

func ClearCacheAllHandler(context *gin.Context) {
	keys, _, err := redisCache.NewRedisCache().Scan(0, "*", constants.ScanCountMax)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": http.StatusInternalServerError,
		})
	}
	for i := range keys {
		redisCache.NewRedisCache().Del(keys[i])
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
	})
}
