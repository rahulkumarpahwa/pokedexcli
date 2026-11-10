package cache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Entry map[string]CacheEntry
	Mu    sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		Entry: make(map[string]CacheEntry),
	}
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	c.Entry[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
	c.Mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	val, ok := c.Entry[key]
	c.Mu.Unlock()
	if ok {
		return val.Val, ok
	}
	return nil, false
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		cache.Mu.Lock()

		for key, entryVal := range cache.Entry {

			if entryVal.CreatedAt.Add(interval).Before(time.Now()) {
				delete(cache.Entry, key)
			}
		}
		cache.Mu.Unlock()

	}

}
