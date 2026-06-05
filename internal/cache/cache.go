package cache

import (
	"sync"
	"time"
)

type Cache struct {
	mu    sync.RWMutex
	items map[string]*item
}

type item struct {
	data      []byte
	expiresAt time.Time
}

func New() *Cache {
	return &Cache{
		items: make(map[string]*item),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	itm, ok := c.items[key]
	if !ok {
		return nil, false
	}

	if time.Now().After(itm.expiresAt) {
		return nil, false
	}

	return itm.data, true
}

func (c *Cache) Set(key string, data []byte, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = &item{
		data:      data,
		expiresAt: time.Now().Add(ttl),
	}
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}
