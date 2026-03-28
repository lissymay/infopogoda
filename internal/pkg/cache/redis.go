package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
	ttl    time.Duration
}

func NewRedisCache(ttl time.Duration, addr, password string, db int) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	ctx := context.Background()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к Redis: %w", err)
	}

	return &RedisCache{
		client: client,
		ctx:    ctx,
		ttl:    ttl,
	}, nil
}

func (r *RedisCache) Get(key string) (interface{}, bool) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return nil, false
	}

	// Пробуем как число
	var temp float64
	if err := json.Unmarshal([]byte(val), &temp); err == nil {
		return float32(temp), true
	}

	// Пробуем как объект (для обратной совместимости)
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(val), &result); err == nil {
		if t, ok := result["temperature_2m"]; ok {
			if tf, ok := t.(float64); ok {
				return float32(tf), true
			}
		}
	}

	return nil, false
}

func (r *RedisCache) Set(key string, value interface{}, ttl time.Duration) {
	if ttl == 0 {
		ttl = r.ttl
	}

	data, err := json.Marshal(value)
	if err != nil {
		return
	}

	r.client.Set(r.ctx, key, data, ttl)
}

func (r *RedisCache) Delete(key string) {
	r.client.Del(r.ctx, key)
}

func (r *RedisCache) Clear() {
	iter := r.client.Scan(r.ctx, 0, "weather:*", 0).Iterator()
	for iter.Next(r.ctx) {
		r.client.Del(r.ctx, iter.Val())
	}
}

func (r *RedisCache) Close() error {
	return r.client.Close()
}
