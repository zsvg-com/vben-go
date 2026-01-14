package cache

import (
	"fmt"
	"time"
	"vben/app/admin/model/constants"
	"vben/pkg/cache/redisCache"
)

type RedisStore struct {
}

// Set 实现设置captcha的方法
func (r RedisStore) Set(id string, value string) error {
	key := constants.CaptchaCodesKey + id
	err := redisCache.NewRedisCache().Put(key, value, time.Minute*3)
	return err
}

// Get 实现获取captcha的方法
func (r RedisStore) Get(id string, clear bool) string {
	key := constants.CaptchaCodesKey + id
	val, err := redisCache.NewRedisCache().Get(key)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		//clear为true，验证通过，删除这个验证码
		_, err := redisCache.NewRedisCache().Del(key)
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

// Verify 实现验证captcha的方法
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
