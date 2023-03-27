package icache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

type LazyCacheItem struct {
	name   string
	handle *cache.Cache
}

var lazyPool sync.Map

func GetLazyCache(cacheName string, cacheTime time.Duration) *LazyCacheItem {
	if value, ok := lazyPool.Load(cacheName); ok {
		return value.(*LazyCacheItem)
	}
	c := cache.New(cacheTime, cacheTime)

	item := &LazyCacheItem{
		name:   cacheName,
		handle: c,
	}
	lazyPool.Store(cacheName, item)

	return item
}

func (c *LazyCacheItem) Set(key string, value interface{}, haveExpire bool) {
	if haveExpire {
		c.handle.Set(key, value, cache.DefaultExpiration)
	} else {
		c.handle.Set(key, value, cache.NoExpiration)
	}
}

func (c *LazyCacheItem) Get(key string) (interface{}, bool) {
	return c.handle.Get(key)
}

func (c *LazyCacheItem) Del(key string) {
	c.handle.Delete(key)
}
