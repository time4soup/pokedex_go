package pokecache

import (
	"sync"
	"time"
)

// individual data for one cache of time and data
type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

// storage of multiple caches and mutex for read/write safety
type Cache struct {
	Data  map[string]CacheEntry
	Mutex sync.Mutex
}

// safely adds new cache at current time to given cache
func (c *Cache) Add(key string, val []byte) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	c.Data[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

// safely gets CacheEntry for given key and returns data and bool of if data was found
func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	entry, ok := c.Data[key]
	if !ok {
		return nil, false
	}
	return entry.Val, true
}

// continuously loops and deletes any entries in given cache older than interval
// loops once every given interval length
func (c *Cache) reapLoop(interval time.Duration) {
	for {
		c.Mutex.Lock()
		for key, entry := range c.Data {
			if time.Since(entry.CreatedAt) > interval {
				delete(c.Data, key)
			}
		}
		c.Mutex.Unlock()

		time.Sleep(interval)
	}
}

// returns newly created empty cache and initializes reaping
func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Data: map[string]CacheEntry{},
	}
	go cache.reapLoop(interval)
	return cache
}
