package ristrettoCache

import (
	"errors"
	"time"
	"vben/pkg/ristretto"
)

type ristrettoCache struct {
}

func NewRistrettoCache() *ristrettoCache {
	return &ristrettoCache{}
}

func (r ristrettoCache) Put(key string, value string, ttl time.Duration) error {
	var isCache = true
	if ttl == 0 {
		isCache = ristretto.Client().Set(key, value, 1)
	} else {
		isCache = ristretto.Client().SetWithTTL(key, value, 1, ttl)
	}
	if isCache {
		return nil
	}
	return errors.New("ristretto is not user")
}

func (r ristrettoCache) Get(key string) (string, error) {
	result, isCache := ristretto.Client().Get(key)
	if isCache {
		return result.(string), nil
	}
	return "", errors.New("ristretto is not user")
}

func (r ristrettoCache) Del(key string) (string, error) {
	ristretto.Client().Del(key)
	return "", nil
}

func (r ristrettoCache) Clear() (string, error) {
	ristretto.Client().Clear()
	return "", nil
}
