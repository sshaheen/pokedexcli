package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
	done     chan struct{}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	b := make([]byte, len(val))
	copy(b, val)
	c.entries[key] = cacheEntry{createdAt: time.Now(), val: b}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.entries[key]
	c.mu.Unlock()

	if !ok {
		return nil, false
	}

	val := make([]byte, len(entry.val))
	copy(val, entry.val)
	return val, true
}

func (c *Cache) reapOnce() {
	c.mu.Lock()
	defer c.mu.Unlock()
	cutoff := time.Now().Add(-c.interval)
	for key, entry := range c.entries {
		if entry.createdAt.Before(cutoff) {
			delete(c.entries, key)
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
		done:     make(chan struct{}),
	}

	ticker := time.NewTicker(interval)

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				c.reapOnce()
			case <-c.done:
				return
			}
		}
	}()

	return c
}

func (c *Cache) Close() {
	close(c.done)
}
