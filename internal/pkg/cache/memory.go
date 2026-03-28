package cache

import (
	"sync"
	"time"
)

type item struct {
	value      interface{}
	expiration int64
}

func (i *item) isExpired() bool {
	if i.expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > i.expiration
}

type MemoryCache struct {
	items map[string]item
	mu    sync.RWMutex
	ttl   time.Duration
}

func NewMemoryCache(defaultTTL time.Duration) *MemoryCache {
	cache := &MemoryCache{
		items: make(map[string]item),
		ttl:   defaultTTL,
	}

	// Запускаем горутину для очистки просроченных записей
	go cache.cleanup()

	return cache
}

func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	if item.isExpired() {
		go c.Delete(key)
		return nil, false
	}

	return item.value, true
}

func (c *MemoryCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var expiration int64
	if ttl == 0 {
		ttl = c.ttl
	}

	if ttl > 0 {
		expiration = time.Now().Add(ttl).UnixNano()
	}

	c.items[key] = item{
		value:      value,
		expiration: expiration,
	}
}

func (c *MemoryCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func (c *MemoryCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]item)
}

// cleanup удаляет просроченные записи
func (c *MemoryCache) cleanup() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		c.mu.Lock()
		for key, item := range c.items {
			if item.isExpired() {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}
