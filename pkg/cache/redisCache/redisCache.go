package redisCache

import (
	"context"
	"time"
	"vben/pkg/redis"
)

type redisCache struct {
}

func NewRedisCache() *redisCache {
	return &redisCache{}
}

func (r redisCache) Put(key string, value string, ttl time.Duration) error {
	_, err := redis.Client().Set(context.TODO(), key, value, ttl).Result()
	return err
}

func (r redisCache) Get(key string) (string, error) {
	return redis.Client().Get(context.TODO(), key).Result()
}

func (r redisCache) Scan(cursor uint64, match string, count int64) (keys []string, c uint64, err error) {
	return redis.Client().Scan(context.TODO(), cursor, match, count).Result()
}

func (r redisCache) Del(key string) (string, error) {
	_, error := redis.Client().Del(context.TODO(), key).Result()
	return "", error
}

func (r redisCache) Clear() (string, error) {
	return "", nil
}

func (r redisCache) Execute(script string, keys []string, args ...interface{}) (interface{}, error) {
	return redis.Client().Eval(context.TODO(), script, keys, args).Result()
}
