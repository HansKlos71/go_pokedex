package pokeapi

import "time"

type Cache struct {
	entries map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for key, entry := range c.entries {
		now := time.Now()
		if now.Sub(entry.createdAt) > 10*time.Minute {
			delete(c.entries, key)
		}
	}
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
