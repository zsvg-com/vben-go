package ristretto

import (
	"github.com/dgraph-io/ristretto"
	"sync"
)

type connect struct {
	client *ristretto.Cache
}

var once = sync.Once{}

var _connect *connect

func connectRistretto() {

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // key 跟踪频率为（10M）
		MaxCost:     1 << 30, // 缓存的最大成本 (1GB)。
		BufferItems: 64,      // 每个 Get buffer的 key 数。
	})
	if err != nil {
		panic(err)
	}

	_connect = &connect{
		client: cache,
	}
}

func Client() *ristretto.Cache {

	if _connect == nil {
		once.Do(func() {
			connectRistretto()
		})
	}

	return _connect.client

}
