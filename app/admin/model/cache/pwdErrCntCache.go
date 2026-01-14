package cache

import (
	"strconv"
	"time"
	"vben/app/admin/model/constants"
	"vben/config"
	"vben/pkg/cache/redisCache"
)

// GetPasswordTryCount	获取输入次数
func GetPasswordTryCount(username string) int {
	countStr, err := redisCache.NewRedisCache().Get(constants.PwdErrCntCacheKey + username)
	if err != nil || len(countStr) <= 0 {
		return 0
	}
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0
	}
	return count
}

// SetPasswordTryCount 次数+1
func SetPasswordTryCount(username string, count int) {
	redisCache.NewRedisCache().Put(constants.PwdErrCntCacheKey+username, strconv.Itoa(count), time.Duration(config.UserPassword.LockTime)*time.Minute)
}

// DeletePasswordTryCount 删除
func DeletePasswordTryCount(username string) {
	redisCache.NewRedisCache().Del(constants.PwdErrCntCacheKey + username)
}
