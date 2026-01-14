package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"vben/app/admin/model/constants"
	"vben/app/core/utils/R"
	"vben/pkg/cache/redisCache"
)

// 定义Lua限流脚本
const script = `  
        local key = KEYS[1]  
        local count = tonumber(ARGV[1])  
        local time = tonumber(ARGV[2])  
        local current = redis.call('get', key)  
        if current and tonumber(current) > count then  
            return tonumber(current)  
        end  
        current = redis.call('incr', key)  
        if tonumber(current) == 1 then  
            redis.call('expire', key, time)  
        end  
        return tonumber(current)  
    `

type LimitType int

const (
	DEFAULT LimitType = iota //限流类型 默认
	Ip                       //根据ip进行限流
)

// RateLimiterMiddleware 限流中间件 每 intervalS 秒 只能请求count次
func RateLimiterMiddleware(intervalS int64, count int64, limitType LimitType) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		combineKey := getCombineKey(limitType, ctx)
		number, err := redisCache.NewRedisCache().Execute(script, []string{combineKey}, count, intervalS)
		if err != nil {
			return
		}
		if number == nil || number.(int64) > count {
			ctx.JSON(http.StatusOK, R.ReturnFailMsg("访问过于频繁，请稍候再试"))
			ctx.Abort()
			return
		}
		log.Printf("限制请求:%v,当前请求:%v,缓存key:%v\n", count, number, combineKey)
	}
}

func getCombineKey(limitType LimitType, ctx *gin.Context) string {
	requestURL := ctx.Request.URL.String()
	method := ctx.Request.Method
	key := constants.RateLimitCacheKey + strings.ReplaceAll(requestURL[1:], "/", "-") + method
	if limitType == Ip {
		key = key + ctx.ClientIP()
	}
	return key
}
