package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mutex        *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cacheEntries: make(map[string]cacheEntry),
		mutex:        &sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}
