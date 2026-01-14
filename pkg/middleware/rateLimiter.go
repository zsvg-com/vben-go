package middleware

import "github.com/gin-gonic/gin"

// 限流
func RateLimiter(l LimiterIface) gin.HandlerFunc {
	return func(context *gin.Context) {
		key := l.Key(context) + context.ClientIP()
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				return
			}
		} else {
			l.AddBucketsByUri(key, 3, 100, 100)
			context.Next()
		}
		context.Next()
	}
}
