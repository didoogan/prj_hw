package services

import store "hw14/internal/repositories"

type Cache struct {
	repository store.CacheRepository
}

func NewCacheService(r store.CacheRepository) *Cache {
	return &Cache{repository: r}
}

func (c *Cache) Get(key string) (value string, exists bool, error error) {
	return c.repository.Get(key)
}

func (c *Cache) Set(key, value string, ttl int) error {
	return c.repository.Set(key, value, ttl)
}
