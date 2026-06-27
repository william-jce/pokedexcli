package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           sync.Mutex
	interval     time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	if key == "" {
		return
	}
	c.mu.Lock()
	c.cacheEntries[key] = cacheEntry{createdAt: time.Now(), val: val}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	found := false
	var entry []byte
	c.mu.Lock()
	if val, ok := c.cacheEntries[key]; ok {
		found = true
		entry = val.val
	}
	c.mu.Unlock()
	return entry, found
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		for entry := range c.cacheEntries {
			if time.Since(c.cacheEntries[entry].createdAt) > c.interval {
				c.mu.Lock()
				delete(c.cacheEntries, entry)
				c.mu.Unlock()
			}
		}
	}
}
