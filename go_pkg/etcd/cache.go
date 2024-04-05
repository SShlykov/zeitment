package etcd

import "sync"

type Cache struct {
	mu     sync.Mutex
	values map[string]string
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Put(key, value string) error {
	c.mu.Lock()
	c.values[key] = value
	c.mu.Unlock()
	return nil
}

func (c *Cache) Get(key string) (string, error) {
	return c.values[key], nil
}

func (c *Cache) Update(values map[string]string) error {
	c.mu.Lock()
	c.values = values
	c.mu.Unlock()

	return nil
}
