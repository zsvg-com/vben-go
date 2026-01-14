package cache

import (
	"time"
	"vben/pkg/cache/fileCache"
	"vben/pkg/cache/redisCache"
	"vben/pkg/cache/ristrettoCache"
)

/**
c := cache.Cache("file")
	c.Put("113", "23", 1*time.Minute)
	result, _ := c.Get("113")
	fmt.Println(result)
*/

type CacheContract interface {
	Put(key string, value string, ttl time.Duration) error
	Get(key string) (string, error)
	Del(key string) (string, error)
	Clear() (string, error)
}

func Cache(driver string) CacheContract {
	switch driver {
	case "redis":
		return redisCache.NewRedisCache()
	case "file":
		return fileCache.NewFileCache()
	case "ristretto":
		return ristrettoCache.NewRistrettoCache()
	}
	return redisCache.NewRedisCache()
}
