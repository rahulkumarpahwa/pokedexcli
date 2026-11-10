package cache

import (
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Entry map[string][]byte
}

func NewCache() *Cache {
	cache := Cache{
		Entry: make(map[string][]byte),
	}
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Entry[key] = val
}

func (c *Cache) Get(key string) ([]byte, bool) {
	val, ok := c.Entry[key]
	if ok {
		return val, ok
	}
	return nil, false
}

