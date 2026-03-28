package cache

import (
	"context"
	"time"
)

// Cache интерфейс для всех кэшей
type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, ttl time.Duration)
	Delete(key string)
	Clear()
}

// RedisCacheInterface отдельный интерфейс для Redis специфичных методов
type RedisCacheInterface interface {
	Cache
	Ping(ctx context.Context) error
	Close() error
}
