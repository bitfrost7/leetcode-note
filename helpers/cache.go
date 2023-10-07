package helpers

import (
	"fmt"
	"sync"
)

type Cache struct {
	mu     sync.Mutex
	result map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		result: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.result[key]
	return v, ok
}
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.result[key] = value
}

func WithCache(cache *Cache, fn func([]any) any) func(...any) any {
	return func(args ...any) any {
		key := fmt.Sprintf("%v", args)
		if v, ok := cache.Get(key); ok {
			return v
		}
		v := fn(args)
		cache.Set(key, v)
		return v
	}
}
