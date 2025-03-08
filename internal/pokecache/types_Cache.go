package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu			sync.RWMutex
	entries		map[string]cacheEntry
	interval 	time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}

	go cache.reapLoop() // Start the reap loop as a goroutine
	return cache
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		<-ticker.C  // Wait for the next tick

		c.mu.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, key) // Remove expired entries
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	timeNow := time.Now()
	newEntry := cacheEntry{
		createdAt:  timeNow,
		val:		val,
	}

	c.entries[key] = newEntry
}

// Accepts a key and returns the cache value and a success flag
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}