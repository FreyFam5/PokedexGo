package pokecache

import (
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheEntries[key] = newEntry
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) (val []byte, found bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, exists := c.cacheEntries[key]
	if !exists {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	for {
		t := <-ticker.C
		c.mutex.Lock()
		for key, entry := range c.cacheEntries {
			if t.After(entry.createdAt.Add(duration)) {
				delete(c.cacheEntries, key)
			}
		}
		c.mutex.Unlock()
	}
}
